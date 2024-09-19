package fiskalhrgo

// CryptoBinary ...
type CryptoBinary string

// Signature ...
type Signature *SignatureType

// SignatureType ...
type SignatureType struct {
	IdAttr           string              `xml:"Id,attr,omitempty"`
	DsSignedInfo     *SignedInfoType     `xml:"ds:SignedInfo"`
	DsSignatureValue *SignatureValueType `xml:"ds:SignatureValue"`
	DsKeyInfo        *KeyInfoType        `xml:"ds:KeyInfo"`
	DsObject         []*ObjectType       `xml:"ds:Object"`
}

// SignatureValue ...
type SignatureValue *SignatureValueType

// SignatureValueType ...
type SignatureValueType struct {
	IdAttr string `xml:"Id,attr,omitempty"`
	Value  string `xml:",chardata"`
}

// SignedInfo ...
type SignedInfo *SignedInfoType

// SignedInfoType ...
type SignedInfoType struct {
	IdAttr                   string                      `xml:"Id,attr,omitempty"`
	DsCanonicalizationMethod *CanonicalizationMethodType `xml:"ds:CanonicalizationMethod"`
	DsSignatureMethod        *SignatureMethodType        `xml:"ds:SignatureMethod"`
	DsReference              []*ReferenceType            `xml:"ds:Reference"`
}

// CanonicalizationMethod ...
type CanonicalizationMethod *CanonicalizationMethodType

// CanonicalizationMethodType ...
type CanonicalizationMethodType struct {
	AlgorithmAttr string `xml:"Algorithm,attr"`
}

// SignatureMethod ...
type SignatureMethod *SignatureMethodType

// SignatureMethodType ...
type SignatureMethodType struct {
	AlgorithmAttr    string `xml:"Algorithm,attr"`
	HMACOutputLength int    `xml:"HMACOutputLength"`
}

// Reference ...
type Reference *ReferenceType

// ReferenceType ...
type ReferenceType struct {
	IdAttr         string            `xml:"Id,attr,omitempty"`
	URIAttr        string            `xml:"URI,attr,omitempty"`
	TypeAttr       string            `xml:"Type,attr,omitempty"`
	DsTransforms   *TransformsType   `xml:"ds:Transforms"`
	DsDigestMethod *DigestMethodType `xml:"ds:DigestMethod"`
	DsDigestValue  *DigestValueType  `xml:"ds:DigestValue"`
}

// Transforms ...
type Transforms *TransformsType

// TransformsType ...
type TransformsType struct {
	DsTransform []*TransformType `xml:"ds:Transform"`
}

// Transform ...
type Transform *TransformType

// TransformType ...
type TransformType struct {
	AlgorithmAttr string   `xml:"Algorithm,attr"`
	XPath         []string `xml:"XPath"`
}

// DigestMethod ...
type DigestMethod *DigestMethodType

// DigestMethodType ...
type DigestMethodType struct {
	AlgorithmAttr string `xml:"Algorithm,attr"`
}

// DigestValue ...
type DigestValue string

// DigestValueType ...
type DigestValueType string

// KeyInfo ...
type KeyInfo *KeyInfoType

// KeyInfoType ...
type KeyInfoType struct {
	IdAttr            string                 `xml:"Id,attr,omitempty"`
	DsKeyName         []string               `xml:"ds:KeyName"`
	DsKeyValue        []*KeyValueType        `xml:"ds:KeyValue"`
	DsRetrievalMethod []*RetrievalMethodType `xml:"ds:RetrievalMethod"`
	DsX509Data        []*X509DataType        `xml:"ds:X509Data"`
	DsPGPData         []*PGPDataType         `xml:"ds:PGPData"`
	DsSPKIData        []*SPKIDataType        `xml:"ds:SPKIData"`
	DsMgmtData        []string               `xml:"ds:MgmtData"`
}

// KeyName ...
type KeyName string

// MgmtData ...
type MgmtData string

// KeyValue ...
type KeyValue *KeyValueType

// KeyValueType ...
type KeyValueType struct {
	DsDSAKeyValue *DSAKeyValueType `xml:"ds:DSAKeyValue"`
	DsRSAKeyValue *RSAKeyValueType `xml:"ds:RSAKeyValue"`
}

// RetrievalMethod ...
type RetrievalMethod *RetrievalMethodType

// RetrievalMethodType ...
type RetrievalMethodType struct {
	URIAttr      string          `xml:"URI,attr,omitempty"`
	TypeAttr     string          `xml:"Type,attr,omitempty"`
	DsTransforms *TransformsType `xml:"ds:Transforms"`
}

// X509Data ...
type X509Data *X509DataType

// X509DataType ...
type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial"`
	X509SKI          string                `xml:"X509SKI"`
	X509SubjectName  string                `xml:"X509SubjectName"`
	X509Certificate  string                `xml:"X509Certificate"`
	X509CRL          string                `xml:"X509CRL"`
}

// X509IssuerSerialType ...
type X509IssuerSerialType struct {
	X509IssuerName   string `xml:"X509IssuerName"`
	X509SerialNumber int    `xml:"X509SerialNumber"`
}

// PGPData ...
type PGPData *PGPDataType

// PGPDataType ...
type PGPDataType struct {
	PGPKeyID     string `xml:"PGPKeyID"`
	PGPKeyPacket string `xml:"PGPKeyPacket"`
}

// SPKIData ...
type SPKIData *SPKIDataType

// SPKIDataType ...
type SPKIDataType struct {
	SPKISexp string `xml:"SPKISexp"`
}

// Object ...
type Object *ObjectType

// ObjectType ...
type ObjectType struct {
	IdAttr       string `xml:"Id,attr,omitempty"`
	MimeTypeAttr string `xml:"MimeType,attr,omitempty"`
	EncodingAttr string `xml:"Encoding,attr,omitempty"`
}

// Manifest ...
type Manifest *ManifestType

// ManifestType ...
type ManifestType struct {
	IdAttr      string           `xml:"Id,attr,omitempty"`
	DsReference []*ReferenceType `xml:"ds:Reference"`
}

// SignatureProperties ...
type SignatureProperties *SignaturePropertiesType

// SignaturePropertiesType ...
type SignaturePropertiesType struct {
	IdAttr              string                   `xml:"Id,attr,omitempty"`
	DsSignatureProperty []*SignaturePropertyType `xml:"ds:SignatureProperty"`
}

// SignatureProperty ...
type SignatureProperty *SignaturePropertyType

// SignaturePropertyType ...
type SignaturePropertyType struct {
	TargetAttr string `xml:"Target,attr"`
	IdAttr     string `xml:"Id,attr,omitempty"`
}

// HMACOutputLengthType ...
type HMACOutputLengthType int

// DSAKeyValue ...
type DSAKeyValue *DSAKeyValueType

// DSAKeyValueType ...
type DSAKeyValueType struct {
	P           string `xml:"P"`
	Q           string `xml:"Q"`
	G           string `xml:"G"`
	Y           string `xml:"Y"`
	J           string `xml:"J"`
	Seed        string `xml:"Seed"`
	PgenCounter string `xml:"PgenCounter"`
}

// RSAKeyValue ...
type RSAKeyValue *RSAKeyValueType

// RSAKeyValueType ...
type RSAKeyValueType struct {
	Modulus  string `xml:"Modulus"`
	Exponent string `xml:"Exponent"`
}
