#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject, Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, Py_GE
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr
from libcpp.optional cimport optional as __optional
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from libcpp.utility cimport move as cmove
from cpython cimport bool as pbool
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
from thrift.py3.types import _IsSet as _fbthrift_IsSet
from thrift.py3.types cimport make_unique
cimport thrift.py3.types
cimport thrift.py3.exceptions
from thrift.python.std_libcpp cimport sv_to_str as __sv_to_str, string_view as __cstring_view
from thrift.py3.types cimport (
    cSetOp as __cSetOp,
    richcmp as __richcmp,
    set_op as __set_op,
    setcmp as __setcmp,
    list_index as __list_index,
    list_count as __list_count,
    list_slice as __list_slice,
    list_getitem as __list_getitem,
    set_iter as __set_iter,
    map_iter as __map_iter,
    map_contains as __map_contains,
    map_getitem as __map_getitem,
    reference_shared_ptr as __reference_shared_ptr,
    get_field_name_by_index as __get_field_name_by_index,
    reset_field as __reset_field,
    translate_cpp_enum_to_python,
    SetMetaClass as __SetMetaClass,
    const_pointer_cast,
    constant_shared_ptr,
    NOTSET as __NOTSET,
    EnumData as __EnumData,
    EnumFlagsData as __EnumFlagsData,
    UnionTypeEnumData as __UnionTypeEnumData,
    createEnumDataForUnionType as __createEnumDataForUnionType,
)
cimport thrift.py3.serializer as serializer
from thrift.python.protocol cimport Protocol as __Protocol
import folly.iobuf as _fbthrift_iobuf
from folly.optional cimport cOptional
from folly.memory cimport to_shared_ptr as __to_shared_ptr
from folly.range cimport Range as __cRange

import sys
from collections.abc import Sequence, Set, Mapping, Iterable
import weakref as __weakref
import builtins as _builtins
import asyncio
from folly.coro cimport bridgeCoroTaskWith



cdef object get_types_reflection():
    import importlib
    return importlib.import_module(
        "module.types_reflection"
    )

@__cython.auto_pickle(False)
cdef class Foo(thrift.py3.types.Struct):
    def __init__(Foo self, **kwargs):
        self._cpp_obj = make_shared[cFoo]()
        self._fields_setter = _fbthrift_types_fields.__Foo_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(Foo self, **kwargs):
        if not kwargs:
            return self
        cdef Foo __fbthrift_inst = Foo.__new__(Foo)
        __fbthrift_inst._cpp_obj = make_shared[cFoo](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__Foo_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("Foo", {
          "value": deref(self._cpp_obj).value_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cFoo] cpp_obj):
        __fbthrift_inst = <Foo>Foo.__new__(Foo)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline value_impl(self):
        if not deref(self._cpp_obj).value_ref().has_value():
            return None

        return deref(self._cpp_obj).value_ref().value_unchecked()

    @property
    def value(self):
        return self.value_impl()


    def __hash__(Foo self):
        return super().__hash__()

    def __repr__(Foo self):
        return super().__repr__()

    def __str__(Foo self):
        return super().__str__()


    def __copy__(Foo self):
        cdef shared_ptr[cFoo] cpp_obj = make_shared[cFoo](
            deref(self._cpp_obj)
        )
        return Foo._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cFoo](
            self._cpp_obj,
            (<Foo>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__Foo()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cFoo].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.Foo"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cFoo](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 1

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(Foo self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cFoo](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(Foo self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cFoo]()
        with nogil:
            needed = serializer.cdeserialize[cFoo](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.Foo, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.Foo, self)

cdef class ClientBufferedStream__i32(ClientBufferedStream):

    @staticmethod
    cdef _fbthrift_create(cClientBufferedStream[cint32_t]& c_obj, __RpcOptions rpc_options):
        __fbthrift_inst = ClientBufferedStream__i32(rpc_options)
        __fbthrift_inst._gen = make_unique[cClientBufferedStreamWrapper[cint32_t]](c_obj)
        return __fbthrift_inst

    @staticmethod
    cdef void callback(
        cFollyTry[__cOptional[cint32_t]]&& result,
        PyObject* userdata,
    ) noexcept:
        cdef __cOptional[cint32_t] opt_val
        cdef cint32_t _value
        stream, pyfuture, rpc_options = <object> userdata
        if result.hasException():
            pyfuture.set_exception(
                thrift.py3.exceptions.create_py_exception(result.exception(), <__RpcOptions>rpc_options)
            )
        else:
            opt_val = result.value()
            if opt_val.has_value():
                try:
                    _value = opt_val.value()
                    pyfuture.set_result(_value)
                except Exception as ex:
                    pyfuture.set_exception(ex.with_traceback(None))
            else:
                pyfuture.set_exception(StopAsyncIteration())

    def __anext__(self):
        __loop = asyncio.get_event_loop()
        __future = __loop.create_future()
        # to avoid "Future exception was never retrieved" error at SIGINT
        __future.add_done_callback(lambda x: x.exception())
        __userdata = (self, __future, self._rpc_options)
        bridgeCoroTaskWith[__cOptional[cint32_t]](
            self._executor,
            deref(self._gen).getNext(),
            ClientBufferedStream__i32.callback,
            <PyObject *>__userdata,
        )
        return asyncio.shield(__future)

cdef class ServerStream__i32(ServerStream):
    pass

