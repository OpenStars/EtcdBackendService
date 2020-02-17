namespace go openstars.pubprofile

struct ProfileData {
  1: string pubkey;
  2: string displayName;
  3: i64 dob;
  4: i64 gender;
  5: string introText;
  6: string avatar;
  7: string imgBackground;
  8: string phone;
  9: string education;
  10: string work;
  11: i64 relationship;
  12: string accommodation;
  13: string linkFB;
  14: string linkGGPlus;
  15: string linkInstagram;
  16: map<string,string> extend;
  17: list<string> image;
  18:i64 lastModified;

}

service PubProfileService {
  ProfileData GetProfileByPubkey(1:string pubkey);
	ProfileData GetProfileByUID(1: i64 uid);
}

