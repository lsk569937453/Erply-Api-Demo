package task

import (
	"context"
	"erply-api/client"
	"erply-api/dao"
	"erply-api/log"
	"erply-api/util"
	"time"

	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New(cron.WithSeconds())
	taskCron := "0/10 * * * * ? "
	// taskCron := "0/5 * * * * ? "
	_, err := c.AddFunc(taskCron, func() {
		pullData()
	})
	if err != nil {
		log.Error("AddFunc error:", err.Error())
		return
	}

	c.Start()
}

func pullData() {
	defer util.ReturnError()
	err := doReq()
	if err != nil {
		log.Error("pullData error", err)
	}
}

func doReq() error {
	customerCli := client.ApiClient.CustomerManager

	reqArray := []map[string]interface{}{
		{
			"recordsOnPage": "100",
			"pageNo":        "1",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	bulkResp, err := customerCli.GetCustomersBulk(ctx, reqArray, map[string]string{})
	if err != nil {
		return err
	}
	resList := bulkResp.BulkItems

	if len(resList) > 0 {
		customers := resList[0].Customers
		if len(customers) > 0 {
			err = UpdateData(customers)
			return err
		}
	}
	return nil
}

func UpdateData(customers []customers.Customer) error {
	for _, item := range customers {
		customerDao := generateDao(item)
		err := UpdateOrInsert(customerDao)
		if err != nil {
			log.Error("UpdateData error", err)
		}
	}
	return nil
}

func UpdateOrInsert(customerDao *dao.CustomerDao) error {
	var count int64
	dao.CronDb.Model(&dao.CustomerDao{}).Where("customer_id = ?", customerDao.CustomerID).Count(&count)
	if count > 0 {
		data := make(map[string]interface{})
		data["group_id"] = customerDao.GroupID
		data["full_name"] = customerDao.FullName
		data["email"] = customerDao.Email
		data["reference_number"] = customerDao.ReferenceNumber
		data["payment_days"] = customerDao.PaymentDays
		err := dao.CronDb.Model(dao.CustomerDao{}).Where("customer_id=?", customerDao.CustomerID).Updates(data).Error
		return err
	} else {
		err := dao.CronDb.Create(customerDao).Error
		return err
	}
}

func generateDao(item customers.Customer) *dao.CustomerDao {
	customerDao := &dao.CustomerDao{
		PayerID:                 item.PayerID,
		CustomerID:              item.CustomerID,
		FullName:                item.FullName,
		CompanyName:             item.CompanyName,
		FirstName:               item.FirstName,
		LastName:                item.LastName,
		GroupID:                 item.GroupID,
		EDI:                     item.EDI,
		GLN:                     item.GLN,
		IsPOSDefaultCustomer:    item.IsPOSDefaultCustomer,
		CountryID:               item.CountryID,
		Phone:                   item.Phone,
		EInvoiceEmail:           item.EInvoiceEmail,
		Email:                   item.Email,
		Fax:                     item.Fax,
		Code:                    item.Code,
		ReferenceNumber:         item.ReferenceNumber,
		VatNumber:               item.VatNumber,
		BankName:                item.BankName,
		BankAccountNumber:       item.BankAccountNumber,
		BankIBAN:                item.BankIBAN,
		BankSWIFT:               item.BankSWIFT,
		PaymentDays:             item.PaymentDays,
		Notes:                   item.Notes,
		LastModified:            item.LastModified,
		CustomerType:            item.CustomerType,
		Address:                 item.Address,
		Street:                  item.Street,
		Address2:                item.Address2,
		City:                    item.City,
		PostalCode:              item.PostalCode,
		Country:                 item.Country,
		State:                   item.State,
		Credit:                  item.Credit,
		CompanyTypeID:           item.CompanyTypeID,
		PersonTitleID:           item.PersonTitleID,
		EmailEnabled:            item.EmailEnabled,
		MailEnabled:             item.MailEnabled,
		EInvoiceEnabled:         item.EInvoiceEnabled,
		FlagStatus:              item.FlagStatus,
		OperatorIdentifier:      item.OperatorIdentifier,
		Gender:                  item.Gender,
		GroupName:               item.GroupName,
		Mobile:                  item.Mobile,
		Birthday:                item.Birthday,
		IntegrationCode:         item.IntegrationCode,
		ColorStatus:             item.ColorStatus,
		FactoringContractNumber: item.FactoringContractNumber,
		Image:                   item.Image,
		TwitterID:               item.TwitterID,
		FacebookName:            item.FacebookName,
		CreditCardLastNumbers:   item.CreditCardLastNumbers,
		EuCustomerType:          item.EuCustomerType,
		CustomerCardNumber:      item.CustomerCardNumber,
		LastModifierUsername:    item.LastModifierUsername,
		DefaultAssociationName:  item.DefaultAssociationName,
		DefaultProfessionalName: item.DefaultProfessionalName,
		TaxExempt:               item.TaxExempt,
		PaysViaFactoring:        item.PaysViaFactoring,
		SalesBlocked:            item.SalesBlocked,
		RewardPointsDisabled:    item.RewardPointsDisabled,
		CustomerBalanceDisabled: item.CustomerBalanceDisabled,
		PosCouponsDisabled:      item.PosCouponsDisabled,
		EmailOptOut:             item.EmailOptOut,
		ShipGoodsWithWaybills:   item.ShipGoodsWithWaybills,
		DefaultAssociationID:    item.DefaultAssociationID,
		DefaultProfessionalID:   item.DefaultProfessionalID,
	}
	return customerDao
}
