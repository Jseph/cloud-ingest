# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/tasks.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/tasks.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x11proto/tasks.proto\"\xb6\x01\n\x0fTaskFailureType\"\xa2\x01\n\x04Type\x12\n\n\x06UNUSED\x10\x00\x12\x0b\n\x07UNKNOWN\x10\x01\x12\x19\n\x15\x46ILE_MODIFIED_FAILURE\x10\x02\x12\x18\n\x14MD5_MISMATCH_FAILURE\x10\x03\x12\x18\n\x14PRECONDITION_FAILURE\x10\x04\x12\x1a\n\x16\x46ILE_NOT_FOUND_FAILURE\x10\x05\x12\x16\n\x12PERMISSION_FAILURE\x10\x06\"G\n\nTaskStatus\"9\n\x04Type\x12\x0c\n\x08UNQUEUED\x10\x00\x12\n\n\x06QUEUED\x10\x01\x12\n\n\x06\x46\x41ILED\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\"9\n\x08TaskType\"-\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04LIST\x10\x01\x12\x0e\n\nUPLOAD_GCS\x10\x02\x42\x07Z\x05protob\x06proto3')
)



_TASKFAILURETYPE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='TaskFailureType.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNUSED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FILE_MODIFIED_FAILURE', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MD5_MISMATCH_FAILURE', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRECONDITION_FAILURE', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FILE_NOT_FOUND_FAILURE', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PERMISSION_FAILURE', index=6, number=6,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=42,
  serialized_end=204,
)
_sym_db.RegisterEnumDescriptor(_TASKFAILURETYPE_TYPE)

_TASKSTATUS_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='TaskStatus.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNQUEUED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='QUEUED', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAILED', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUCCESS', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=220,
  serialized_end=277,
)
_sym_db.RegisterEnumDescriptor(_TASKSTATUS_TYPE)

_TASKTYPE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='TaskType.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LIST', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UPLOAD_GCS', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=291,
  serialized_end=336,
)
_sym_db.RegisterEnumDescriptor(_TASKTYPE_TYPE)


_TASKFAILURETYPE = _descriptor.Descriptor(
  name='TaskFailureType',
  full_name='TaskFailureType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TASKFAILURETYPE_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=204,
)


_TASKSTATUS = _descriptor.Descriptor(
  name='TaskStatus',
  full_name='TaskStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TASKSTATUS_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=206,
  serialized_end=277,
)


_TASKTYPE = _descriptor.Descriptor(
  name='TaskType',
  full_name='TaskType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TASKTYPE_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=279,
  serialized_end=336,
)

_TASKFAILURETYPE_TYPE.containing_type = _TASKFAILURETYPE
_TASKSTATUS_TYPE.containing_type = _TASKSTATUS
_TASKTYPE_TYPE.containing_type = _TASKTYPE
DESCRIPTOR.message_types_by_name['TaskFailureType'] = _TASKFAILURETYPE
DESCRIPTOR.message_types_by_name['TaskStatus'] = _TASKSTATUS
DESCRIPTOR.message_types_by_name['TaskType'] = _TASKTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TaskFailureType = _reflection.GeneratedProtocolMessageType('TaskFailureType', (_message.Message,), dict(
  DESCRIPTOR = _TASKFAILURETYPE,
  __module__ = 'proto.tasks_pb2'
  # @@protoc_insertion_point(class_scope:TaskFailureType)
  ))
_sym_db.RegisterMessage(TaskFailureType)

TaskStatus = _reflection.GeneratedProtocolMessageType('TaskStatus', (_message.Message,), dict(
  DESCRIPTOR = _TASKSTATUS,
  __module__ = 'proto.tasks_pb2'
  # @@protoc_insertion_point(class_scope:TaskStatus)
  ))
_sym_db.RegisterMessage(TaskStatus)

TaskType = _reflection.GeneratedProtocolMessageType('TaskType', (_message.Message,), dict(
  DESCRIPTOR = _TASKTYPE,
  __module__ = 'proto.tasks_pb2'
  # @@protoc_insertion_point(class_scope:TaskType)
  ))
_sym_db.RegisterMessage(TaskType)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\005proto'))
# @@protoc_insertion_point(module_scope)
