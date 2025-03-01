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


cdef __EnumData __MyEnum_enum_data  = __EnumData._fbthrift_create(thrift.py3.types.createEnumData[cMyEnum](), MyEnum)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __MyEnumMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __MyEnum_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __MyEnum_enum_data.get_all_names()

    def __len__(cls):
        return __MyEnum_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __MyEnum_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class MyEnum(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __MyEnum_enum_data.get_by_name(name)


    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        EnumMetadata[cMyEnum].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.MyEnum"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return python_types.MyEnum(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value


__SetMetaClass(<PyTypeObject*> MyEnum, <PyTypeObject*> __MyEnumMeta)



cdef object get_types_reflection():
    import importlib
    return importlib.import_module(
        "module.types_reflection"
    )

@__cython.auto_pickle(False)
cdef class MyStructNestedAnnotation(thrift.py3.types.Struct):
    def __init__(MyStructNestedAnnotation self, **kwargs):
        self._cpp_obj = make_shared[cMyStructNestedAnnotation]()
        self._fields_setter = _fbthrift_types_fields.__MyStructNestedAnnotation_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(MyStructNestedAnnotation self, **kwargs):
        if not kwargs:
            return self
        cdef MyStructNestedAnnotation __fbthrift_inst = MyStructNestedAnnotation.__new__(MyStructNestedAnnotation)
        __fbthrift_inst._cpp_obj = make_shared[cMyStructNestedAnnotation](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__MyStructNestedAnnotation_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("MyStructNestedAnnotation", {
          "name": deref(self._cpp_obj).name_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cMyStructNestedAnnotation] cpp_obj):
        __fbthrift_inst = <MyStructNestedAnnotation>MyStructNestedAnnotation.__new__(MyStructNestedAnnotation)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline name_impl(self):

        return (<bytes>deref(self._cpp_obj).name_ref().value()).decode('UTF-8')

    @property
    def name(self):
        return self.name_impl()


    def __hash__(MyStructNestedAnnotation self):
        return super().__hash__()

    def __repr__(MyStructNestedAnnotation self):
        return super().__repr__()

    def __str__(MyStructNestedAnnotation self):
        return super().__str__()


    def __copy__(MyStructNestedAnnotation self):
        cdef shared_ptr[cMyStructNestedAnnotation] cpp_obj = make_shared[cMyStructNestedAnnotation](
            deref(self._cpp_obj)
        )
        return MyStructNestedAnnotation._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cMyStructNestedAnnotation](
            self._cpp_obj,
            (<MyStructNestedAnnotation>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__MyStructNestedAnnotation()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cMyStructNestedAnnotation].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.MyStructNestedAnnotation"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cMyStructNestedAnnotation](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 1

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(MyStructNestedAnnotation self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cMyStructNestedAnnotation](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(MyStructNestedAnnotation self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cMyStructNestedAnnotation]()
        with nogil:
            needed = serializer.cdeserialize[cMyStructNestedAnnotation](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.MyStructNestedAnnotation, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.MyStructNestedAnnotation, self)
@__cython.auto_pickle(False)
cdef class SecretStruct(thrift.py3.types.Struct):
    def __init__(SecretStruct self, **kwargs):
        self._cpp_obj = make_shared[cSecretStruct]()
        self._fields_setter = _fbthrift_types_fields.__SecretStruct_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(SecretStruct self, **kwargs):
        if not kwargs:
            return self
        cdef SecretStruct __fbthrift_inst = SecretStruct.__new__(SecretStruct)
        __fbthrift_inst._cpp_obj = make_shared[cSecretStruct](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__SecretStruct_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("SecretStruct", {
          "id": deref(self._cpp_obj).id_ref().has_value(),
          "password": deref(self._cpp_obj).password_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cSecretStruct] cpp_obj):
        __fbthrift_inst = <SecretStruct>SecretStruct.__new__(SecretStruct)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline id_impl(self):

        return deref(self._cpp_obj).id_ref().value()

    @property
    def id(self):
        return self.id_impl()

    cdef inline password_impl(self):

        return (<bytes>deref(self._cpp_obj).password_ref().value()).decode('UTF-8')

    @property
    def password(self):
        return self.password_impl()


    def __hash__(SecretStruct self):
        return super().__hash__()

    def __repr__(SecretStruct self):
        return super().__repr__()

    def __str__(SecretStruct self):
        return super().__str__()


    def __copy__(SecretStruct self):
        cdef shared_ptr[cSecretStruct] cpp_obj = make_shared[cSecretStruct](
            deref(self._cpp_obj)
        )
        return SecretStruct._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cSecretStruct](
            self._cpp_obj,
            (<SecretStruct>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__SecretStruct()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cSecretStruct].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.SecretStruct"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cSecretStruct](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 2

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(SecretStruct self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cSecretStruct](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(SecretStruct self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cSecretStruct]()
        with nogil:
            needed = serializer.cdeserialize[cSecretStruct](buf, self._cpp_obj.get(), proto)
        return needed


    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.SecretStruct, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.SecretStruct, self)
@__cython.auto_pickle(False)
cdef class std_deque_std_string__List__string(thrift.py3.types.List):
    def __init__(self, items=None):
        if isinstance(items, std_deque_std_string__List__string):
            self._cpp_obj = (<std_deque_std_string__List__string> items)._cpp_obj
        else:
            self._cpp_obj = std_deque_std_string__List__string._make_instance(items)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[std_deque_std_string] c_items):
        __fbthrift_inst = <std_deque_std_string__List__string>std_deque_std_string__List__string.__new__(std_deque_std_string__List__string)
        __fbthrift_inst._cpp_obj = cmove(c_items)
        return __fbthrift_inst

    def __copy__(std_deque_std_string__List__string self):
        cdef shared_ptr[std_deque_std_string] cpp_obj = make_shared[std_deque_std_string](
            deref(self._cpp_obj)
        )
        return std_deque_std_string__List__string._fbthrift_create(cmove(cpp_obj))

    def __len__(self):
        return deref(self._cpp_obj).size()

    @staticmethod
    cdef shared_ptr[std_deque_std_string] _make_instance(object items) except *:
        cdef shared_ptr[std_deque_std_string] c_inst = make_shared[std_deque_std_string]()
        if items is not None:
            if isinstance(items, str):
                raise TypeError("If you really want to pass a string into a _typing.Sequence[str] field, explicitly convert it first.")
            for item in items:
                if not isinstance(item, str):
                    raise TypeError(f"{item!r} is not of type str")
                deref(c_inst).push_back(item.encode('UTF-8'))
        return c_inst

    cdef _get_slice(self, slice index_obj):
        cdef int start, stop, step
        start, stop, step = index_obj.indices(deref(self._cpp_obj).size())
        return std_deque_std_string__List__string._fbthrift_create(
            __list_slice[std_deque_std_string](self._cpp_obj, start, stop, step)
        )

    cdef _get_single_item(self, size_t index):
        cdef string citem
        __list_getitem(self._cpp_obj, index, citem)
        return bytes(citem).decode('UTF-8')

    cdef _check_item_type(self, item):
        if not self or item is None:
            return
        if isinstance(item, str):
            return item

    def index(self, item, start=0, stop=None):
        err = ValueError(f'{item} is not in list')
        item = self._check_item_type(item)
        if item is None:
            raise err
        cdef (int, int, int) indices = slice(start, stop).indices(deref(self._cpp_obj).size())
        cdef string citem = item.encode('UTF-8')
        cdef __optional[size_t] found = __list_index[std_deque_std_string](self._cpp_obj, indices[0], indices[1], citem)
        if not found.has_value():
            raise err
        return found.value()

    def count(self, item):
        item = self._check_item_type(item)
        if item is None:
            return 0
        cdef string citem = item.encode('UTF-8')
        return __list_count[std_deque_std_string](self._cpp_obj, citem)

    @staticmethod
    def __get_reflection__():
        return get_types_reflection().get_reflection__std_deque_std_string__List__string()


Sequence.register(std_deque_std_string__List__string)

myStruct = MyStruct._fbthrift_create(constant_shared_ptr(cmyStruct()))
list_string_6884 = std_deque_std_string__List__string
