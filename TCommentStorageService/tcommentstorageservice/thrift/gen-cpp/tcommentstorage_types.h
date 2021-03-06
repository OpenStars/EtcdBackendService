/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef tcommentstorage_TYPES_H
#define tcommentstorage_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>


namespace OpenStars { namespace Common { namespace TCommentStorageService {

struct TErrorCode {
  enum type {
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2,
    EDataExisted = -3
  };
};

extern const std::map<int, const char*> _TErrorCode_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const TErrorCode::type& val);

typedef int64_t TKey;

typedef class TCommentItem TData;

class ActionLink;

class MediaItem;

class TCommentItem;

class TDataResult;

class TListDataResult;

typedef struct _ActionLink__isset {
  _ActionLink__isset() : text(false), href(false) {}
  bool text :1;
  bool href :1;
} _ActionLink__isset;

class ActionLink : public virtual ::apache::thrift::TBase {
 public:

  ActionLink(const ActionLink&);
  ActionLink& operator=(const ActionLink&);
  ActionLink() : text(), href() {
  }

  virtual ~ActionLink() throw();
  std::string text;
  std::string href;

  _ActionLink__isset __isset;

  void __set_text(const std::string& val);

  void __set_href(const std::string& val);

  bool operator == (const ActionLink & rhs) const
  {
    if (__isset.text != rhs.__isset.text)
      return false;
    else if (__isset.text && !(text == rhs.text))
      return false;
    if (__isset.href != rhs.__isset.href)
      return false;
    else if (__isset.href && !(href == rhs.href))
      return false;
    return true;
  }
  bool operator != (const ActionLink &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const ActionLink & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(ActionLink &a, ActionLink &b);

std::ostream& operator<<(std::ostream& out, const ActionLink& obj);

typedef struct _MediaItem__isset {
  _MediaItem__isset() : name(false), mediaType(false), url(false) {}
  bool name :1;
  bool mediaType :1;
  bool url :1;
} _MediaItem__isset;

class MediaItem : public virtual ::apache::thrift::TBase {
 public:

  MediaItem(const MediaItem&);
  MediaItem& operator=(const MediaItem&);
  MediaItem() : name(), mediaType(0), url() {
  }

  virtual ~MediaItem() throw();
  std::string name;
  int64_t mediaType;
  std::string url;

  _MediaItem__isset __isset;

  void __set_name(const std::string& val);

  void __set_mediaType(const int64_t val);

  void __set_url(const std::string& val);

  bool operator == (const MediaItem & rhs) const
  {
    if (!(name == rhs.name))
      return false;
    if (!(mediaType == rhs.mediaType))
      return false;
    if (!(url == rhs.url))
      return false;
    return true;
  }
  bool operator != (const MediaItem &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const MediaItem & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(MediaItem &a, MediaItem &b);

std::ostream& operator<<(std::ostream& out, const MediaItem& obj);

typedef struct _TCommentItem__isset {
  _TCommentItem__isset() : idcomment(false), uid(false), pubkey(false), idpost(false), content(false), actionlinks(false), mediaitem(false), pubkeyTags(false), uidTags(false), timestamps(false), parentcommentid(false) {}
  bool idcomment :1;
  bool uid :1;
  bool pubkey :1;
  bool idpost :1;
  bool content :1;
  bool actionlinks :1;
  bool mediaitem :1;
  bool pubkeyTags :1;
  bool uidTags :1;
  bool timestamps :1;
  bool parentcommentid :1;
} _TCommentItem__isset;

class TCommentItem : public virtual ::apache::thrift::TBase {
 public:

  TCommentItem(const TCommentItem&);
  TCommentItem& operator=(const TCommentItem&);
  TCommentItem() : idcomment(0), uid(0), pubkey(), idpost(0), content(), timestamps(0), parentcommentid(0) {
  }

  virtual ~TCommentItem() throw();
  TKey idcomment;
  int64_t uid;
  std::string pubkey;
  int64_t idpost;
  std::string content;
  std::vector<ActionLink>  actionlinks;
  MediaItem mediaitem;
  std::vector<std::string>  pubkeyTags;
  std::vector<std::string>  uidTags;
  int64_t timestamps;
  int64_t parentcommentid;

  _TCommentItem__isset __isset;

  void __set_idcomment(const TKey val);

  void __set_uid(const int64_t val);

  void __set_pubkey(const std::string& val);

  void __set_idpost(const int64_t val);

  void __set_content(const std::string& val);

  void __set_actionlinks(const std::vector<ActionLink> & val);

  void __set_mediaitem(const MediaItem& val);

  void __set_pubkeyTags(const std::vector<std::string> & val);

  void __set_uidTags(const std::vector<std::string> & val);

  void __set_timestamps(const int64_t val);

  void __set_parentcommentid(const int64_t val);

  bool operator == (const TCommentItem & rhs) const
  {
    if (!(idcomment == rhs.idcomment))
      return false;
    if (!(uid == rhs.uid))
      return false;
    if (!(pubkey == rhs.pubkey))
      return false;
    if (!(idpost == rhs.idpost))
      return false;
    if (!(content == rhs.content))
      return false;
    if (__isset.actionlinks != rhs.__isset.actionlinks)
      return false;
    else if (__isset.actionlinks && !(actionlinks == rhs.actionlinks))
      return false;
    if (__isset.mediaitem != rhs.__isset.mediaitem)
      return false;
    else if (__isset.mediaitem && !(mediaitem == rhs.mediaitem))
      return false;
    if (__isset.pubkeyTags != rhs.__isset.pubkeyTags)
      return false;
    else if (__isset.pubkeyTags && !(pubkeyTags == rhs.pubkeyTags))
      return false;
    if (__isset.uidTags != rhs.__isset.uidTags)
      return false;
    else if (__isset.uidTags && !(uidTags == rhs.uidTags))
      return false;
    if (!(timestamps == rhs.timestamps))
      return false;
    if (__isset.parentcommentid != rhs.__isset.parentcommentid)
      return false;
    else if (__isset.parentcommentid && !(parentcommentid == rhs.parentcommentid))
      return false;
    return true;
  }
  bool operator != (const TCommentItem &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TCommentItem & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TCommentItem &a, TCommentItem &b);

std::ostream& operator<<(std::ostream& out, const TCommentItem& obj);

typedef struct _TDataResult__isset {
  _TDataResult__isset() : errorCode(false), data(false) {}
  bool errorCode :1;
  bool data :1;
} _TDataResult__isset;

class TDataResult : public virtual ::apache::thrift::TBase {
 public:

  TDataResult(const TDataResult&);
  TDataResult& operator=(const TDataResult&);
  TDataResult() : errorCode((TErrorCode::type)0) {
  }

  virtual ~TDataResult() throw();
  TErrorCode::type errorCode;
  TCommentItem data;

  _TDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_data(const TCommentItem& val);

  bool operator == (const TDataResult & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (__isset.data != rhs.__isset.data)
      return false;
    else if (__isset.data && !(data == rhs.data))
      return false;
    return true;
  }
  bool operator != (const TDataResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TDataResult &a, TDataResult &b);

std::ostream& operator<<(std::ostream& out, const TDataResult& obj);

typedef struct _TListDataResult__isset {
  _TListDataResult__isset() : errorCode(false), listDatas(false) {}
  bool errorCode :1;
  bool listDatas :1;
} _TListDataResult__isset;

class TListDataResult : public virtual ::apache::thrift::TBase {
 public:

  TListDataResult(const TListDataResult&);
  TListDataResult& operator=(const TListDataResult&);
  TListDataResult() : errorCode((TErrorCode::type)0) {
  }

  virtual ~TListDataResult() throw();
  TErrorCode::type errorCode;
  std::vector<TCommentItem>  listDatas;

  _TListDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_listDatas(const std::vector<TCommentItem> & val);

  bool operator == (const TListDataResult & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (!(listDatas == rhs.listDatas))
      return false;
    return true;
  }
  bool operator != (const TListDataResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TListDataResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TListDataResult &a, TListDataResult &b);

std::ostream& operator<<(std::ostream& out, const TListDataResult& obj);

}}} // namespace

#endif
