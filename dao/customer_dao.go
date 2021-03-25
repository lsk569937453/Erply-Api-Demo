package dao

const (
	OperationCollectHoldingFurnaceDefault = 1
	OperationCollectRationFurnace         = 0
	OperationSetTemperature               = 2
	OperationCollectHoldingFurnaceSetTem  = 3

	DEVICE_ONLINE  = 0
	DEVICE_OFFLINE = 1
)

type CustomerDao struct {
	ID int `gorm:"primary_key;auto_increment" json:"id"`

	PayerID int `gorm:"type:varchar(50);not null;index:pay_id_idx;comment:'pay id'" json:"payerID,omitempty"`

	CustomerID              int    `gorm:"type:varchar(50);not null;uniqueIndex:customer_id_idx;comment:'customer id'"  json:"customerID"`
	TypeID                  string `gorm:"type:varchar(50);not null;default:'';index:type_id_idx;comment:'type id'" json:"type_id"`
	FullName                string `gorm:"type:varchar(50);not null;default:'';comment:'full name'" json:"fullName"`
	CompanyName             string `gorm:"type:varchar(50);not null;default:'';comment:'company name'" json:"companyName"`
	FirstName               string `gorm:"type:varchar(20);not null;default:'';comment:'first name'" json:"firstName"`
	LastName                string `gorm:"type:varchar(20);not null;default:'';comment:'Last name'" json:"lastName"`
	GroupID                 int    `gorm:"type:varchar(50);not null;default:0;index:group_id_idx;comment:'group id'" json:"groupID"`
	EDI                     string `gorm:"type:varchar(50);not null;default:'';comment:'Edi'" json:"EDI"`
	GLN                     string `gorm:"type:varchar(50);not null;default:'';comment:'Gln'" json:"GLN"`
	IsPOSDefaultCustomer    int    `gorm:"type:int(10);not null;default:0;index:is_pos_default_customer_idx;comment:'is POSDefault Customer'"  json:"isPOSDefaultCustomer"`
	CountryID               string `gorm:"type:varchar(50);not null;default:0;index:contry_id_idx;comment:'country ID'" json:"countryID"`
	Phone                   string `gorm:"type:varchar(50);not null;default:0;index:phone_idx;comment:'phone'" json:"phone"`
	EInvoiceEmail           string `gorm:"type:varchar(50);not null;default:'';comment:'EInvoice Email'" json:"eInvoiceEmail"`
	Email                   string `gorm:"type:varchar(50);not null;default:'';comment:'EInvoice Email'" json:"email"`
	Fax                     string `gorm:"type:varchar(50);not null;default:'';comment:'Fax'" json:"fax"`
	Code                    string `gorm:"type:varchar(50);not null;default:'';comment:'Code'" json:"code"`
	ReferenceNumber         string `gorm:"type:varchar(50);not null;default:'';comment:'Reference Number'" json:"referenceNumber"`
	VatNumber               string `gorm:"type:varchar(50);not null;default:'';comment:'Vat Number'" json:"vatNumber"`
	BankName                string `gorm:"type:varchar(50);not null;default:'';comment:'Bank Name'" json:"bankName"`
	BankAccountNumber       string `gorm:"type:varchar(50);not null;default:'';comment:'Bank Account Number'" json:"bankAccountNumber"`
	BankIBAN                string `gorm:"type:varchar(50);not null;default:'';comment:'Bank IBan'" json:"bankIBAN"`
	BankSWIFT               string `gorm:"type:varchar(50);not null;default:'';comment:'Bank BankSWIFT'" json:"bankSWIFT"`
	PaymentDays             int    `gorm:"type:int(10);not null;default:0;index:payment_days_idx;comment:'Payment Days'"  json:"paymentDays"`
	Notes                   string `gorm:"type:varchar(50);not null;default:'';comment:'Notes'" json:"notes"`
	LastModified            int    `gorm:"type:int(10);not null;default:0;comment:'Last Modified'" json:"lastModified"`
	CustomerType            string `gorm:"type:varchar(50);not null;default:'';comment:'Customer Type'" json:"customerType"`
	Address                 string `gorm:"type:varchar(50);not null;default:'';comment:'Address'" json:"address"`
	Street                  string `gorm:"type:varchar(50);not null;default:'';comment:'street'" json:"street"`
	Address2                string `gorm:"type:varchar(50);not null;default:'';comment:'address2'" json:"address2"`
	City                    string `gorm:"type:varchar(50);not null;default:'';comment:'city'" json:"city"`
	PostalCode              string `gorm:"type:varchar(50);not null;default:'';comment:'postal code'" json:"postalCode"`
	Country                 string `gorm:"type:varchar(50);not null;default:'';comment:'country'" json:"country"`
	State                   string `gorm:"type:varchar(50);not null;default:'';comment:'state'" json:"state"`
	Credit                  int    `gorm:"type:int(10);not null;default:0;comment:'credit'" json:"credit"`
	CompanyTypeID           int    `gorm:"type:int(10);not null;default:0;comment:'company Type ID'" json:"companyTypeID"`
	PersonTitleID           int    `gorm:"type:int(10);not null;default:0;comment:'Person  Title Id'" json:"personTitleID"`
	EmailEnabled            int    `gorm:"type:int(10);not null;default:0;comment:'Email enabled'" json:"emailEnabled"`
	MailEnabled             int    `gorm:"type:int(10);not null;default:0;comment:'Mail enabled'" json:"mailEnabled"`
	EInvoiceEnabled         int    `gorm:"type:int(10);not null;default:0;comment:'EInvoice enabled'" json:"eInvoiceEnabled"`
	FlagStatus              int    `gorm:"type:int(10);not null;default:0;comment:'Flag Status'" json:"flagStatus"`
	OperatorIdentifier      string `gorm:"type:varchar(50);not null;default:'';comment:'Operator Identifier'" json:"operatorIdentifier"`
	Gender                  string `gorm:"type:varchar(50);not null;default:'';comment:'gender'" json:"gender"`
	GroupName               string `gorm:"type:varchar(50);not null;default:'';comment:'groupName'" json:"groupName"`
	Mobile                  string `gorm:"type:varchar(50);not null;default:'';comment:'mobile'" json:"mobile"`
	Birthday                string `gorm:"type:varchar(50);not null;default:'';comment:'birthday'" json:"birthday"`
	IntegrationCode         string `gorm:"type:varchar(50);not null;default:'';comment:'Integration Code'" json:"integrationCode"`
	ColorStatus             string `gorm:"type:varchar(50);not null;default:'';comment:'Color Status'" json:"colorStatus"`
	FactoringContractNumber string `gorm:"type:varchar(50);not null;default:'';comment:'Factoring Contract Number'" json:"factoringContractNumber"`
	Image                   string `gorm:"type:varchar(50);not null;default:'';comment:'Image'" json:"image"`
	TwitterID               string `gorm:"type:varchar(50);not null;default:'';comment:'Twitter ID'" json:"twitterID"`
	FacebookName            string `gorm:"type:varchar(50);not null;default:'';comment:'Facebook Name'" json:"facebookName"`
	CreditCardLastNumbers   string `gorm:"type:varchar(50);not null;default:'';comment:'Credit Card Last Numbers'" json:"creditCardLastNumbers"`
	EuCustomerType          string `gorm:"type:varchar(50);not null;default:'';comment:'Eu Customer Type'" json:"euCustomerType"`
	CustomerCardNumber      string `gorm:"type:varchar(50);not null;default:'';comment:'Customer Card Number'" json:"customerCardNumber"`
	LastModifierUsername    string `gorm:"type:varchar(50);not null;default:'';comment:'Last Modifier Username'" json:"lastModifierUsername"`
	DefaultAssociationName  string `gorm:"type:varchar(50);not null;default:'';comment:'Default Association Name'" json:"defaultAssociationName"`
	DefaultProfessionalName string `gorm:"type:varchar(50);not null;default:'';comment:'Default Professional Name'" json:"defaultProfessionalName"`
	TaxExempt               int    `gorm:"type:int(10);not null;default:0;comment:'Tax Exempt'" json:"taxExempt"`
	PaysViaFactoring        int    `gorm:"type:int(10);not null;default:0;comment:'Pays Via Factoring'" json:"paysViaFactoring"`
	SalesBlocked            int    `gorm:"type:int(10);not null;default:0;comment:'Sales Blocked'" json:"salesBlocked"`
	RewardPointsDisabled    int    `gorm:"type:int(10);not null;default:0;comment:'Reward Points Disabled'" json:"rewardPointsDisabled"`
	CustomerBalanceDisabled int    `gorm:"type:int(10);not null;default:0;comment:'Customer Balance Disabled'" json:"customerBalanceDisabled"`
	PosCouponsDisabled      int    `gorm:"type:int(10);not null;default:0;comment:'PosCoupons Disabled'" json:"posCouponsDisabled"`
	EmailOptOut             int    `gorm:"type:int(10);not null;default:0;comment:'Email Opt Out'" json:"emailOptOut"`
	ShipGoodsWithWaybills   int    `gorm:"type:int(10);not null;default:0;comment:'Ship Goods WithWay bills'" json:"shipGoodsWithWaybills"`
	DefaultAssociationID    int    `gorm:"type:int(10);not null;default:0;comment:'Default AssociationID'" json:"defaultAssociationID"`
	DefaultProfessionalID   int    `gorm:"type:int(10);not null;default:0;comment:'Default ProfessionalID'" json:"defaultProfessionalID"`

	Timestamp SelfFormatTime `gorm:"type:timestamp(3) on update current_timestamp(3);omitempty;index:timestamp_idx;default:current_timestamp(3);comment:'时间戳'" json:"timestamp"`
}

func (p CustomerDao) TableName() string {
	return "customer"
}
