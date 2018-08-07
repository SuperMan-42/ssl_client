package ssl

const (
	WARNING AlertLevel = 1
	FATAL AlertLevel = 2
)

const (
	CLOSE_NOTIFY AlertDescription = 0
	UNEXPECTED_MESSAGE = 10
	BAD_RECORD_MAC = 20
	DECOMPRESSION_FAILURE = 30
	HANDSHAKE_FAILURE = 40
	NO_CERTIFICATE = 41
	BAD_CERTIFICATE = 42
	UNSUPPORTED_CERTIFICATE = 43
	CERTIFICATE_REVOKED = 44
	CERTIFICATE_EXPIRED = 45
	CERTIFICATE_UNKNOWN = 46
	ILLEGAL_PARAMETER = 47
)

const (
	TLS_NULL_WITH_NULL_NULL CipherSuite = 0x0000

	/*
	The following CipherSuite definitions require that the server provide
	an RSA certificate that can be used for key exchange.  The server may
	request any signature-capable certificate in the certificate request
	message.  */

	TLS_RSA_WITH_NULL_MD5           = 0x0001
	TLS_RSA_WITH_NULL_SHA           = 0x0002
	TLS_RSA_EXPORT_WITH_RC4_40_MD5  = 0x0003
	TLS_RSA_WITH_NULL_SHA256        = 0x003B
	TLS_RSA_WITH_RC4_128_MD5        = 0x0004
	TLS_RSA_WITH_RC4_128_SHA        = 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA   = 0x000A
	TLS_RSA_WITH_AES_128_CBC_SHA    = 0x002F
	TLS_RSA_WITH_AES_256_CBC_SHA    = 0x0035
	TLS_RSA_WITH_AES_128_CBC_SHA256 = 0x003C
	TLS_RSA_WITH_AES_256_CBC_SHA256 = 0x003D

	/*
	The following cipher suite definitions are used for server-
	authenticated (and optionally client-authenticated) Diffie-Hellman.
	DH denotes cipher suites in which the server's certificate contains
	the Diffie-Hellman parameters signed by the certificate authority
	(CA).  DHE denotes ephemeral Diffie-Hellman, where the Diffie-Hellman
	parameters are signed by a signature-capable certificate, which has
	been signed by the CA.  The signing algorithm used by the server is
	specified after the DHE component of the CipherSuite name.  The
	server can request any signature-capable certificate from the client
	for client authentication, or it may request a Diffie-Hellman
	certificate.  Any Diffie-Hellman certificate provided by the client
	must use the parameters (group and generator) described by the
	server.
	*/

	TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA    = 0x000D
	TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA    = 0x0010
	TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA   = 0x0013
	TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA   = 0x0016
	TLS_DH_DSS_WITH_AES_128_CBC_SHA     = 0x0030
	TLS_DH_RSA_WITH_AES_128_CBC_SHA     = 0x0031
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA    = 0x0032
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA    = 0x0033
	TLS_DH_DSS_WITH_AES_256_CBC_SHA     = 0x0036
	TLS_DH_RSA_WITH_AES_256_CBC_SHA     = 0x0037
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA    = 0x0038
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA    = 0x0039
	TLS_DH_DSS_WITH_AES_128_CBC_SHA256  = 0x003E
	TLS_DH_RSA_WITH_AES_128_CBC_SHA256  = 0x003F
	TLS_DHE_DSS_WITH_AES_128_CBC_SHA256 = 0x0040
	TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 = 0x0067
	TLS_DH_DSS_WITH_AES_256_CBC_SHA256  = 0x0068
	TLS_DH_RSA_WITH_AES_256_CBC_SHA256  = 0x0069
	TLS_DHE_DSS_WITH_AES_256_CBC_SHA256 = 0x006A
	TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 = 0x006B

	/*
	The following cipher suites are used for completely anonymous
	Diffie-Hellman communications in which neither party is
	authenticated.  Note that this mode is vulnerable to man-in-the-
	middle attacks.  Using this mode therefore is of limited use: These
	cipher suites MUST NOT be used by TLS 1.2 implementations unless the
	application layer has specifically requested to allow anonymous key
	exchange.  (Anonymous key exchange may sometimes be acceptable, for
	example, to support opportunistic encryption when no set-up for
	authentication is in place, or when TLS is used as part of more
	complex security protocols that have other means to ensure
	authentication.)
	 */

	TLS_DH_anon_WITH_RC4_128_MD5        = 0x0018
	TLS_DH_anon_WITH_3DES_EDE_CBC_SHA   = 0x001B
	TLS_DH_anon_WITH_AES_128_CBC_SHA    = 0x0034
	TLS_DH_anon_WITH_AES_256_CBC_SHA    = 0x003A
	TLS_DH_anon_WITH_AES_128_CBC_SHA256 = 0x006C
	TLS_DH_anon_WITH_AES_256_CBC_SHA256 = 0x006D

	TLS_AES_128_GCM_SHA256       = 0x1301
	TLS_AES_128_GCM_SHA384       = 0x1302
	TLS_CHACHA20_POLY1305_SHA256 = 0x1303
	TLS_AES_128_CCM_SHA256       = 0x1304
	TLS_AES_128_CCM_8_SHA256     = 0x1305

	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       = 0xc02c
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384         = 0xc030
	TLS_DHE_DSS_WITH_AES_256_GCM_SHA384           = 0x00a3
	TLS_DHE_RSA_WITH_AES_256_GCM_SHA384           = 0x009f
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 = 0xcca9
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   = 0xcca8
	TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256     = 0xccaa
	TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8            = 0xc0af
	TLS_ECDHE_ECDSA_WITH_AES_256_CCM              = 0xc0ad
	TLS_DHE_RSA_WITH_AES_256_CCM_8                = 0xc0a3
	TLS_DHE_RSA_WITH_AES_256_CCM                  = 0xc09f
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384       = 0xc024
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384         = 0xc028
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384  = 0xc073
	TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384    = 0xc077
	TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256      = 0x00c4
	TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256      = 0x00c3
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          = 0xc00a
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA            = 0xc014
	TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA         = 0x0088
	TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA         = 0x0087
	TLS_RSA_WITH_AES_256_GCM_SHA384               = 0x009d
	TLS_RSA_WITH_AES_256_CCM_8                    = 0xc0a1
	TLS_RSA_WITH_AES_256_CCM                      = 0xc09d
	TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256          = 0x00c0
	TLS_RSA_WITH_CAMELLIA_256_CBC_SHA             = 0x0084
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       = 0xc02b
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256         = 0xc02f
	TLS_DHE_DSS_WITH_AES_128_GCM_SHA256           = 0x00a2
	TLS_DHE_RSA_WITH_AES_128_GCM_SHA256           = 0x009e
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8            = 0xc0ae
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM              = 0xc0ac
	TLS_DHE_RSA_WITH_AES_128_CCM_8                = 0xc0a2
	TLS_DHE_RSA_WITH_AES_128_CCM                  = 0xc09e
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       = 0xc023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256         = 0xc027
	TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256  = 0xc072
	TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256    = 0xc076
	TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256      = 0x00be
	TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256      = 0x00bd
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          = 0xc009
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA            = 0xc013
	TLS_DHE_RSA_WITH_SEED_CBC_SHA                 = 0x009a
	TLS_DHE_DSS_WITH_SEED_CBC_SHA                 = 0x0099
	TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA         = 0x0045
	TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA         = 0x0044
	TLS_RSA_WITH_AES_128_GCM_SHA256               = 0x009c
	TLS_RSA_WITH_AES_128_CCM_8                    = 0xc0a0
	TLS_RSA_WITH_AES_128_CCM                      = 0xc09c
	TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256          = 0x00ba
	TLS_RSA_WITH_SEED_CBC_SHA                     = 0x0096
	TLS_RSA_WITH_CAMELLIA_128_CBC_SHA             = 0x0041
	TLS_RSA_WITH_IDEA_CBC_SHA                     = 0x0007
	TLS_EMPTY_RENEGOTIATION_INFO_SCSV             = 0x00ff
)

const (
	NULL_COMPRESSION CompressionMethod = 0
)

const (
	CHANGE_CIPHER_SPEC ContentType = 20
	ALERT                          = 21
	HANDSHAKE                      = 22
	APPLICATION_DATA               = 23
)

const (
	HELLO_REQUEST       HandshakeType = 0
	CLIENT_HELLO                      = 1
	SERVER_HELLO                      = 2
	CERTIFICATE                       = 11
	SERVER_KEY_EXCHANGE               = 12
	CERTIFICATE_REQUEST               = 13
	SERVER_HELLO_DONE                 = 14
	CERTIFICATE_VERIFY                = 15
	CLIENT_KEY_EXCHANGE               = 16
	FINISHED                          = 20
)