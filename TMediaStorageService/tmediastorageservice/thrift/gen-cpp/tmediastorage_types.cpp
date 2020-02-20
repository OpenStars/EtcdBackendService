/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "tmediastorage_types.h"

#include <algorithm>
#include <ostream>

#include <thrift/TToString.h>

namespace OpenStars { namespace Common { namespace TMediaStorageService {

int _kTErrorCodeValues[] = {
  TErrorCode::EGood,
  TErrorCode::ENotFound,
  TErrorCode::EUnknown,
  TErrorCode::EDataExisted
};
const char* _kTErrorCodeNames[] = {
  "EGood",
  "ENotFound",
  "EUnknown",
  "EDataExisted"
};
const std::map<int, const char*> _TErrorCode_VALUES_TO_NAMES(::apache::thrift::TEnumIterator(4, _kTErrorCodeValues, _kTErrorCodeNames), ::apache::thrift::TEnumIterator(-1, NULL, NULL));

std::ostream& operator<<(std::ostream& out, const TErrorCode::type& val) {
  std::map<int, const char*>::const_iterator it = _TErrorCode_VALUES_TO_NAMES.find(val);
  if (it != _TErrorCode_VALUES_TO_NAMES.end()) {
    out << it->second;
  } else {
    out << static_cast<int>(val);
  }
  return out;
}


TMediaItem::~TMediaItem() throw() {
}


void TMediaItem::__set_name(const std::string& val) {
  this->name = val;
}

void TMediaItem::__set_mediaType(const int64_t val) {
  this->mediaType = val;
}

void TMediaItem::__set_url(const std::string& val) {
  this->url = val;
}

void TMediaItem::__set_idmedia(const int64_t val) {
  this->idmedia = val;
}

void TMediaItem::__set_idpost(const int64_t val) {
  this->idpost = val;
}

void TMediaItem::__set_timestamps(const int64_t val) {
  this->timestamps = val;
}

void TMediaItem::__set_extend(const std::string& val) {
  this->extend = val;
}

void TMediaItem::__set_mapExtend(const std::map<std::string, std::string> & val) {
  this->mapExtend = val;
}
std::ostream& operator<<(std::ostream& out, const TMediaItem& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TMediaItem::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->name);
          this->__isset.name = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->mediaType);
          this->__isset.mediaType = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 3:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->url);
          this->__isset.url = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 4:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->idmedia);
          this->__isset.idmedia = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 5:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->idpost);
          this->__isset.idpost = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 6:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->timestamps);
          this->__isset.timestamps = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 7:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->extend);
          this->__isset.extend = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 8:
        if (ftype == ::apache::thrift::protocol::T_MAP) {
          {
            this->mapExtend.clear();
            uint32_t _size0;
            ::apache::thrift::protocol::TType _ktype1;
            ::apache::thrift::protocol::TType _vtype2;
            xfer += iprot->readMapBegin(_ktype1, _vtype2, _size0);
            uint32_t _i4;
            for (_i4 = 0; _i4 < _size0; ++_i4)
            {
              std::string _key5;
              xfer += iprot->readString(_key5);
              std::string& _val6 = this->mapExtend[_key5];
              xfer += iprot->readString(_val6);
            }
            xfer += iprot->readMapEnd();
          }
          this->__isset.mapExtend = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TMediaItem::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TMediaItem");

  xfer += oprot->writeFieldBegin("name", ::apache::thrift::protocol::T_STRING, 1);
  xfer += oprot->writeString(this->name);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("mediaType", ::apache::thrift::protocol::T_I64, 2);
  xfer += oprot->writeI64(this->mediaType);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("url", ::apache::thrift::protocol::T_STRING, 3);
  xfer += oprot->writeString(this->url);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("idmedia", ::apache::thrift::protocol::T_I64, 4);
  xfer += oprot->writeI64(this->idmedia);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("idpost", ::apache::thrift::protocol::T_I64, 5);
  xfer += oprot->writeI64(this->idpost);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("timestamps", ::apache::thrift::protocol::T_I64, 6);
  xfer += oprot->writeI64(this->timestamps);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("extend", ::apache::thrift::protocol::T_STRING, 7);
  xfer += oprot->writeString(this->extend);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("mapExtend", ::apache::thrift::protocol::T_MAP, 8);
  {
    xfer += oprot->writeMapBegin(::apache::thrift::protocol::T_STRING, ::apache::thrift::protocol::T_STRING, static_cast<uint32_t>(this->mapExtend.size()));
    std::map<std::string, std::string> ::const_iterator _iter7;
    for (_iter7 = this->mapExtend.begin(); _iter7 != this->mapExtend.end(); ++_iter7)
    {
      xfer += oprot->writeString(_iter7->first);
      xfer += oprot->writeString(_iter7->second);
    }
    xfer += oprot->writeMapEnd();
  }
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TMediaItem &a, TMediaItem &b) {
  using ::std::swap;
  swap(a.name, b.name);
  swap(a.mediaType, b.mediaType);
  swap(a.url, b.url);
  swap(a.idmedia, b.idmedia);
  swap(a.idpost, b.idpost);
  swap(a.timestamps, b.timestamps);
  swap(a.extend, b.extend);
  swap(a.mapExtend, b.mapExtend);
  swap(a.__isset, b.__isset);
}

TMediaItem::TMediaItem(const TMediaItem& other8) {
  name = other8.name;
  mediaType = other8.mediaType;
  url = other8.url;
  idmedia = other8.idmedia;
  idpost = other8.idpost;
  timestamps = other8.timestamps;
  extend = other8.extend;
  mapExtend = other8.mapExtend;
  __isset = other8.__isset;
}
TMediaItem& TMediaItem::operator=(const TMediaItem& other9) {
  name = other9.name;
  mediaType = other9.mediaType;
  url = other9.url;
  idmedia = other9.idmedia;
  idpost = other9.idpost;
  timestamps = other9.timestamps;
  extend = other9.extend;
  mapExtend = other9.mapExtend;
  __isset = other9.__isset;
  return *this;
}
void TMediaItem::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TMediaItem(";
  out << "name=" << to_string(name);
  out << ", " << "mediaType=" << to_string(mediaType);
  out << ", " << "url=" << to_string(url);
  out << ", " << "idmedia=" << to_string(idmedia);
  out << ", " << "idpost=" << to_string(idpost);
  out << ", " << "timestamps=" << to_string(timestamps);
  out << ", " << "extend=" << to_string(extend);
  out << ", " << "mapExtend=" << to_string(mapExtend);
  out << ")";
}


TDataResult::~TDataResult() throw() {
}


void TDataResult::__set_errorCode(const TErrorCode::type val) {
  this->errorCode = val;
}

void TDataResult::__set_data(const TMediaItem& val) {
  this->data = val;
__isset.data = true;
}
std::ostream& operator<<(std::ostream& out, const TDataResult& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TDataResult::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          int32_t ecast10;
          xfer += iprot->readI32(ecast10);
          this->errorCode = (TErrorCode::type)ecast10;
          this->__isset.errorCode = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_STRUCT) {
          xfer += this->data.read(iprot);
          this->__isset.data = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TDataResult::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TDataResult");

  xfer += oprot->writeFieldBegin("errorCode", ::apache::thrift::protocol::T_I32, 1);
  xfer += oprot->writeI32((int32_t)this->errorCode);
  xfer += oprot->writeFieldEnd();

  if (this->__isset.data) {
    xfer += oprot->writeFieldBegin("data", ::apache::thrift::protocol::T_STRUCT, 2);
    xfer += this->data.write(oprot);
    xfer += oprot->writeFieldEnd();
  }
  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TDataResult &a, TDataResult &b) {
  using ::std::swap;
  swap(a.errorCode, b.errorCode);
  swap(a.data, b.data);
  swap(a.__isset, b.__isset);
}

TDataResult::TDataResult(const TDataResult& other11) {
  errorCode = other11.errorCode;
  data = other11.data;
  __isset = other11.__isset;
}
TDataResult& TDataResult::operator=(const TDataResult& other12) {
  errorCode = other12.errorCode;
  data = other12.data;
  __isset = other12.__isset;
  return *this;
}
void TDataResult::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TDataResult(";
  out << "errorCode=" << to_string(errorCode);
  out << ", " << "data="; (__isset.data ? (out << to_string(data)) : (out << "<null>"));
  out << ")";
}


TListDataResult::~TListDataResult() throw() {
}


void TListDataResult::__set_errorCode(const TErrorCode::type val) {
  this->errorCode = val;
}

void TListDataResult::__set_listDatas(const std::vector<TMediaItem> & val) {
  this->listDatas = val;
}
std::ostream& operator<<(std::ostream& out, const TListDataResult& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TListDataResult::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          int32_t ecast13;
          xfer += iprot->readI32(ecast13);
          this->errorCode = (TErrorCode::type)ecast13;
          this->__isset.errorCode = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_LIST) {
          {
            this->listDatas.clear();
            uint32_t _size14;
            ::apache::thrift::protocol::TType _etype17;
            xfer += iprot->readListBegin(_etype17, _size14);
            this->listDatas.resize(_size14);
            uint32_t _i18;
            for (_i18 = 0; _i18 < _size14; ++_i18)
            {
              xfer += this->listDatas[_i18].read(iprot);
            }
            xfer += iprot->readListEnd();
          }
          this->__isset.listDatas = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TListDataResult::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TListDataResult");

  xfer += oprot->writeFieldBegin("errorCode", ::apache::thrift::protocol::T_I32, 1);
  xfer += oprot->writeI32((int32_t)this->errorCode);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("listDatas", ::apache::thrift::protocol::T_LIST, 2);
  {
    xfer += oprot->writeListBegin(::apache::thrift::protocol::T_STRUCT, static_cast<uint32_t>(this->listDatas.size()));
    std::vector<TMediaItem> ::const_iterator _iter19;
    for (_iter19 = this->listDatas.begin(); _iter19 != this->listDatas.end(); ++_iter19)
    {
      xfer += (*_iter19).write(oprot);
    }
    xfer += oprot->writeListEnd();
  }
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TListDataResult &a, TListDataResult &b) {
  using ::std::swap;
  swap(a.errorCode, b.errorCode);
  swap(a.listDatas, b.listDatas);
  swap(a.__isset, b.__isset);
}

TListDataResult::TListDataResult(const TListDataResult& other20) {
  errorCode = other20.errorCode;
  listDatas = other20.listDatas;
  __isset = other20.__isset;
}
TListDataResult& TListDataResult::operator=(const TListDataResult& other21) {
  errorCode = other21.errorCode;
  listDatas = other21.listDatas;
  __isset = other21.__isset;
  return *this;
}
void TListDataResult::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TListDataResult(";
  out << "errorCode=" << to_string(errorCode);
  out << ", " << "listDatas=" << to_string(listDatas);
  out << ")";
}

}}} // namespace