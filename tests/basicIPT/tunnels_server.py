#============LICENSE_START=============================================================================================================
# Copyright (C) 2020 AT&T Intellectual Property. All rights reserved.
#===================================================================
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#============LICENSE_END===============================================================================================================

"""The Python implementation of the WB Application Offload tunnel table server."""

from concurrent import futures
import time
import logging
import socket
import struct
import sys

import grpc
import google.protobuf.timestamp_pb2

#from grpc_status import rpc_status

from google.protobuf import any_pb2
#from google.rpc import code_pb2, status_pb2, error_details_pb2

import tunneloffload_pb2
import tunneloffload_pb2_grpc


class TunnelValidationExcp(Exception):
    pass

class AddTunnelErrors:
   '''
   '''
   def __init__(self):
       self._addTunnelErrors= list()
   def addTunnelErrorMembers(self, tunnelResponseError):
       self._addTunnelErrors.append(tunnelResponseError)
   def __iter__(self):
       ''' Returns the Iterator object '''
       return AddTunnelErrorsIterator(self)

class AddTunnelErrorsIterator:
   ''' Iterator class '''
   def __init__(self, addTunnelErrors_list):
       # AddTunnelErrors object reference
       self._addTunnelErrors_list = addTunnelErrors_list
       # member variable to keep track of current index
       self._index = 0
   def __next__(self):
       ''''Returns the next value from team object's lists '''
       if self._index < (len(self._addTunnelErrors_list._addTunnelErrors) ) :
           result = (self._addTunnelErrors_list._addTunnelErrors[self._index])
           self._index +=1
           return result
       # End of Iteration
       raise StopIteration


class TunnelValidators(object):
    def __init__(self, tunnel):
        self.tunnel = tunnel
        self.match_criteria = self.tunnel.match_criteria

    @staticmethod
    def validate_ipv6_pair(ipv6_pair):
        if len(ipv6_pair.sourceIp) != 32:
            raise TunnelValidationExcp("IPV6 Source IP Isn't valid")

        if len(ipv6_pair.destinationIp) != 32:
            raise TunnelValidationExcp("IPV6 Destination IP Isn't valid")

    @staticmethod
    def validate_mac_pair(mac_pair):
        if len(mac_pair.sourceMac) != 12:
            raise TunnelValidationExcp("Source MAC isn't valid in geneveEncap")
        
        if len(mac_pair.destinationMac) != 12:
            raise TunnelValidationExcp("Destination MAC isn't valid in geneveEncap")

    def check_geneve_encap_validity(self, geneve_encap):
        # Checking MAC's validity
        self.validate_mac_pair(geneve_encap.innerMacPair)

        if geneve_encap.vni > (2**24 -1):
            raise TunnelValidationExcp("VNI is more than 24 bit, invalid")

        outer_ip =  geneve_encap.WhichOneof('ip')
        if not outer_ip:
            raise TunnelValidationExcp("No outer IP encapsulation found on geneve encap, invalid")

        if outer_ip == 'outerIpv6Pair':
            self.validate_ipv6_pair(geneve_encap.outerIpv6Pair)

        
    def check_geneve_decap_validity(self, geneve_decap):
        if self.match_criteria.WhichOneof('match') != 'geneveMatch':
            raise TunnelValidationExcp("GENEVE Decap action without geneve match isnt valid")

    def check_geneve_validity(self, geneve_params):
        geneve_encap_decap = geneve_params.WhichOneof('encap_decap')

        if geneve_encap_decap == 'geneveEncap':
            self.check_geneve_encap_validity(getattr(geneve_params, geneve_encap_decap))
        else:
            self.check_geneve_decap_validity(getattr(geneve_params, geneve_encap_decap))

    def check_ipsec_tunnel_validity(self, ipsec_params):
        ipsec_mode = ipsec_params.WhichOneof("ipsec")

        if ipsec_mode == "ipsecDec":
            ipsec_params = ipsec_params.ipsecDec
        else:
            ipsec_params = ipsec_params.ipsecEnc 

        if ipsec_params.encryptionType not in [tunneloffload_pb2._AES256GCM8, tunneloffload_pb2._AES256GCM12, tunneloffload_pb2._AES256GCM16]:
            raise TunnelValidationExcp("Not encryption type set on tunnel")

        if len(ipsec_params.encryptionKey) != 64: # Checking the key length
            raise TunnelValidationExcp("Key used for IPSec isnt in right size of 256-bit")

        # Enc specific checks
        if ipsec_mode == "ipsecEnc":
            if ipsec_params.SPI == 0:
                raise TunnelValidationExcp("SPI isn't set on IPSec Enc Request")

    def check_tunnel_validity(self):

        TUNNEL_TO_VALIDITY_DICT = {
        'geneve': self.check_geneve_validity,
        'ipsecTunnel': self.check_ipsec_tunnel_validity,   
        }

        tunnel_type = self.tunnel.WhichOneof("tunnel") 
        if tunnel_type is None:
            raise TunnelValidationExcp("No tunnel type passed to tunnel")

        if tunnel_type not in TUNNEL_TO_VALIDITY_DICT:
            raise TunnelValidationExcp("Internel problem in server, no validation to tunnel")

        TUNNEL_TO_VALIDITY_DICT[tunnel_type](getattr(self.tunnel, tunnel_type))

    
class ipTunnelServiceServicer(tunneloffload_pb2_grpc.ipTunnelServiceServicer):
    """Provides methods that implement functionality of tunnel table server."""
    """ rpc createIpTunnels(ipTunnel) returns (createIpTunnelResponse) {} """


    def __init__(self):
        """do some init stuff"""

    def createIpTunnel(self, request_iterator, context):
        tunnelErrors_value=AddTunnelErrors()
        for request in request_iterator:
            print("############ Create Tunnel ##################")
            print("TunnelId:", request.tunnelId)
            print("Match.ipv4Match.sourceIp", socket.inet_ntop(socket.AF_INET, request.match_criteria.ipv4Match.sourceIp.to_bytes(4,byteorder=sys.byteorder)))
            print("Match.ipv4Match.destinationIp", socket.inet_ntop(socket.AF_INET, request.match_criteria.ipv4Match.destinationIp.to_bytes(4,byteorder=sys.byteorder)))
            print("nextAction:" , request.nextAction)
            print("Checking validity of the request")
            tunnel_validator = TunnelValidators(request)
            tunnel_validator.check_tunnel_validity()
            

        return tunneloffload_pb2.createIpTunnelResponse(requestStatus=tunneloffload_pb2._TUNNEL_ACCEPTED)






def tunnelServe():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    tunneloffload_pb2_grpc.add_ipTunnelServiceServicer_to_server(
        ipTunnelServiceServicer(), server)
    with open('ssl/server.key', 'rb') as f:
         private_key = f.read()
    with open('ssl/server.crt', 'rb') as f:
         certificate_chain = f.read()
    server_credentials = grpc.ssl_server_credentials( ( (private_key, certificate_chain), ) )
    server.add_secure_port('localhost:3443', server_credentials)
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    tunnelServe()