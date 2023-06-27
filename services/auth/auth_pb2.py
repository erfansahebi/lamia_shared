# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/auth/auth.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18services/auth/auth.proto\x12\x04\x61uth\"`\n\nUserStruct\x12\n\n\x02id\x18\x01 \x01(\t\x12\x12\n\nfirst_name\x18\x02 \x01(\t\x12\x11\n\tlast_name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x10\n\x08password\x18\x05 \x01(\t\"U\n\x16\x41uthenticationResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\x10.auth.UserStruct\x12\x1b\n\x13\x61uthorization_token\x18\x02 \x01(\t\"1\n\x0fRegisterRequest\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\x10.auth.UserStruct\"/\n\x0cLoginRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\",\n\rLogoutRequest\x12\x1b\n\x13\x61uthorization_token\x18\x01 \x01(\t\"\x10\n\x0eLogoutResponse\"2\n\x13\x41uthenticateRequest\x12\x1b\n\x13\x61uthorization_token\x18\x01 \x01(\t\"\"\n\x14\x41uthenticateResponse\x12\n\n\x02id\x18\x01 \x01(\t\"!\n\x0eGetUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"1\n\x0fGetUserResponse\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\x10.auth.UserStruct2\xc7\x02\n\x0b\x41uthService\x12\x41\n\x08Register\x12\x15.auth.RegisterRequest\x1a\x1c.auth.AuthenticationResponse\"\x00\x12;\n\x05Login\x12\x12.auth.LoginRequest\x1a\x1c.auth.AuthenticationResponse\"\x00\x12\x35\n\x06Logout\x12\x13.auth.LogoutRequest\x1a\x14.auth.LogoutResponse\"\x00\x12G\n\x0c\x41uthenticate\x12\x19.auth.AuthenticateRequest\x1a\x1a.auth.AuthenticateResponse\"\x00\x12\x38\n\x07GetUser\x12\x14.auth.GetUserRequest\x1a\x15.auth.GetUserResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.auth.auth_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _globals['_USERSTRUCT']._serialized_start=34
  _globals['_USERSTRUCT']._serialized_end=130
  _globals['_AUTHENTICATIONRESPONSE']._serialized_start=132
  _globals['_AUTHENTICATIONRESPONSE']._serialized_end=217
  _globals['_REGISTERREQUEST']._serialized_start=219
  _globals['_REGISTERREQUEST']._serialized_end=268
  _globals['_LOGINREQUEST']._serialized_start=270
  _globals['_LOGINREQUEST']._serialized_end=317
  _globals['_LOGOUTREQUEST']._serialized_start=319
  _globals['_LOGOUTREQUEST']._serialized_end=363
  _globals['_LOGOUTRESPONSE']._serialized_start=365
  _globals['_LOGOUTRESPONSE']._serialized_end=381
  _globals['_AUTHENTICATEREQUEST']._serialized_start=383
  _globals['_AUTHENTICATEREQUEST']._serialized_end=433
  _globals['_AUTHENTICATERESPONSE']._serialized_start=435
  _globals['_AUTHENTICATERESPONSE']._serialized_end=469
  _globals['_GETUSERREQUEST']._serialized_start=471
  _globals['_GETUSERREQUEST']._serialized_end=504
  _globals['_GETUSERRESPONSE']._serialized_start=506
  _globals['_GETUSERRESPONSE']._serialized_end=555
  _globals['_AUTHSERVICE']._serialized_start=558
  _globals['_AUTHSERVICE']._serialized_end=885
# @@protoc_insertion_point(module_scope)