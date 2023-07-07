# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import shop_pb2 as proto_dot_shop_dot_shop__pb2


class ShopServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/shop.ShopService/Create',
                request_serializer=proto_dot_shop_dot_shop__pb2.CreateRequest.SerializeToString,
                response_deserializer=proto_dot_shop_dot_shop__pb2.CreateResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/shop.ShopService/Get',
                request_serializer=proto_dot_shop_dot_shop__pb2.GetRequest.SerializeToString,
                response_deserializer=proto_dot_shop_dot_shop__pb2.GetResponse.FromString,
                )


class ShopServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ShopServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=proto_dot_shop_dot_shop__pb2.CreateRequest.FromString,
                    response_serializer=proto_dot_shop_dot_shop__pb2.CreateResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=proto_dot_shop_dot_shop__pb2.GetRequest.FromString,
                    response_serializer=proto_dot_shop_dot_shop__pb2.GetResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'shop.ShopService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ShopService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shop.ShopService/Create',
            proto_dot_shop_dot_shop__pb2.CreateRequest.SerializeToString,
            proto_dot_shop_dot_shop__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/shop.ShopService/Get',
            proto_dot_shop_dot_shop__pb2.GetRequest.SerializeToString,
            proto_dot_shop_dot_shop__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
