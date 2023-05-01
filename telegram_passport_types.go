package main

type PassportData struct {
	// Array with information about documents and other
	// Telegram Passport elements that was shared with the bot
	Data []EncryptedPassportElement `json:"data"`

	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the
	// same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`

	// File size in bytes
	FileSize int `json:"file_size"`

	// Unix time when the file was uploaded
	FileDate int `json:"file_date"`
}

type EncryptedPassportElement struct {
	// Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`

	// Optional. Base64-encoded encrypted Telegram Passport element data provided by the user,
	// available for “personal_details”, “passport”, “driver_license”, “identity_card”,
	// “internal_passport” and “address” types. Can be decrypted and verified using
	// the accompanying EncryptedCredentials.
	Data string `json:"data"`

	// Optional. User's verified phone number, available only for “phone_number” type
	PhoneNumber string `json:"phone_number"`

	// Optional. User's verified email address, available only for “email” type
	Email string `json:"email"`

	// Optional. Array of encrypted files with documents provided by the user,
	// available for “utility_bill”, “bank_statement”, “rental_agreement”,
	// “passport_registration” and “temporary_registration” types. Files can be decrypted
	// and verified using the accompanying EncryptedCredentials.
	Files []PassportFile `json:"files"`

	// Optional. Encrypted file with the front side of the document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side"`

	// Optional. Encrypted file with the reverse side of the document, provided by the user.
	// Available for “driver_license” and “identity_card”. The file can be decrypted and
	// verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side"`

	// Optional. Encrypted file with the selfie of the user holding a document, provided
	// by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie"`

	// Optional. Array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”,
	// “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and
	// “temporary_registration” types. Files can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	Translation []PassportFile `json:"translation"`

	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	Hash string `json:"hash"`
}

type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and
	// secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`

	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`

	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

type PassportElementErrorDataField struct {
	// Error source, must be data
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the error, one of “personal_details”,
	// “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	Type string `json:"type"`

	// Name of the data field which has the error
	FieldName string `json:"field_name"`

	// Base64-encoded data hash
	DataHash string `json:"data_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorFrontSide struct {
	// Error source, must be front_side
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the front side of the document
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorReverseSide struct {
	// Error source, must be reverse_side
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue,
	// one of “driver_license”, “identity_card”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the reverse side of the document
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorSelfie struct {
	// Error source, must be selfie
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”
	Type string `json:"type"`

	// Base64-encoded hash of the file with the selfie
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorFile struct {
	// Error source, must be file
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorFiles struct {
	// Error source, must be files
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorTranslationFile struct {
	// Error source, must be translation_file
	Source string `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue, one of “passport”,
	// “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”,
	// “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
	// Error source, must be translation_files
	Source string `json:"source"`

	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”,
	// “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	Type string `json:"type"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes"`

	// Error message
	Message string `json:"message"`
}

type PassportElementErrorUnspecified struct {
	// Error source, must be unspecified
	Source string `json:"source"`

	// Type of element of the user's Telegram Passport which has the issue
	Type string `json:"type"`

	// Base64-encoded element hash
	ElementHash string `json:"element_hash"`

	// Error message
	Message string `json:"message"`
}
