https://appspector.com/blog/how-to-improve-messagepack-javascript-parsing-speed-by-2-6-times

msgpack is slower for some things and faster for others in Dart and JS, it seems like wins over JSON aren't really worth the loss of readability.

Go msgpack libraries are 10x faster than the standard lib json. I wonder if easyjson or something similar could provide a good middle-of-the-road option here between performance and readability.

# jenn
Go web service generators

Note to self: Inspired by https://youtu.be/j6ow-UemzBc

### future improvements

// todo filter out fields by namespace
// todo use generated JSON stubs for perf improvements


### Things to look into
- PostgREST
- Stitch
- FeathersJS
- GraphQL
- Meteor
- Hasura
- Apollo
- Slash GraphQL

### Types

#### Primitives
- String
- Float
- Int
- Bool
- List
- Number
- ByteArray
- JSONBlob
- Serial

#### Complex Scalar Types

- FirstLastName
- FullName
- SSN(includeMask: bool)

- Timestamp(format: TimeFormat)
- TimeOfDay
- Datetime
- Timezone

- PostalAddress(useAbbreviations: bool)
- State
- City
- Country
- Zipcode
- ZipcodeLong
- LatLong

- PhoneNumber(usOnly: bool, includePlus: bool, includeMask: bool)
- Email
- Avatar(frameSize: BoxSize, maxSizeMB: Int)
- Password(minLength: Int, maxLength: Int, requireCaps: bool, requireSpecialChars: bool)
- Domain
- URL

- CCNumber
- Currency(supportedCurrencies: list)

- IPv4(allowCIDR: bool)
- IPv6
- Port
- Domain
- SemanticVersion
- FileHash
- File
- JSONWebToken
- RandomString
- RandomInt
- RSA Key
- Image
- Identicon
- Pincode
- CryptoHash (MD5, SHA, etc)

### Full System Model

Apps have a single endpoint

Apps are comprised of the following:

- Functions
    Used to run business logic when events occur and call APIs
- Event Streams
    Used for auditing, as well as Functions and Service-to-Service sharing
- APIs
    CRUD services that provide data, validation and output events. Uni-directional flow of data (API Call -> Event)
- Authenticators
    Microservices that provide authentication and can link into APIs to provide authorization data. These include SMS-OTP, Email-OTL, Google, Facebook, Auth0, Apple, Github, Gitlab, LDAP, ID/Password

Single endpoint, client gets manifest to determine what APIs are available, what environment to use, what version... also if sandbox is available.

Staging environment / Shared Sandbox

Sandbox (spun up on demand, isolated to just a single user)