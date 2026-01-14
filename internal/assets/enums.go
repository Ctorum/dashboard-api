package assets

type Personas string
type Document string
type FileType string
type FileClass string
type FileMimeType string
type AssetStatus string
type AssetType string

const (
	PersonasDebtor     Personas = "DEBTOR"
	PersonasAssignor   Personas = "ASSIGNOR"
	PersonasGuarantor  Personas = "GUARANTOR"
	PersonasSupplier   Personas = "SUPPLIER"
	PersonasIntegrator Personas = "INTEGRATOR"
)

const (
	DocumentPF Document = "PF"
	DocumentPJ Document = "PJ"
)

const (
	FileTypeCCB                 FileType = "CCB"
	FileTypeProofOfEndorsement  FileType = "PROOF_OF_ENDORSEMENT"
	FileTypeProofOfDisbursement FileType = "PROOF_OF_DISBURSEMENT"
	FileTypeDebtorDocument      FileType = "DEBTOR_DOCUMENT"
	FileTypeDuplicate           FileType = "DUPLICATE"
	FileTypeInvoice             FileType = "INVOICE"
	FileTypeAssignmentDocument  FileType = "ASSIGNMENT_DOCUMENT"
)

const (
	FileClassCollateral    FileClass = "COLLATERAL"
	FileClassSupplementary FileClass = "SUPPLEMENTARY"
)

const (
	FileMimeTypePDF   FileMimeType = "PDF"
	FileMimeTypeXML   FileMimeType = "XML"
	FileMimeTypeImage FileMimeType = "IMAGE"
	FileMimeTypeCSV   FileMimeType = "CSV"
	FileMimeTypeJSON  FileMimeType = "JSON"
)

const (
	AssetStatusCreated         AssetStatus = "CREATED"
	AssetStatusValidationError AssetStatus = "VALIDATION_ERROR"
	AssetStatusProcessingError AssetStatus = "PROCESSING_ERROR"
	AssetStatusDiscarded       AssetStatus = "DISCARDED"
	AssetStatusPreApproved     AssetStatus = "PRE_APPROVED"
	AssetStatusApproved        AssetStatus = "APPROVED"
)

const (
	AssetTypeTradeReceivable AssetType = "TRADE_RECEIVABLE"
	AssetTypeTaxInvoice      AssetType = "TAX_INVOICE"
)
