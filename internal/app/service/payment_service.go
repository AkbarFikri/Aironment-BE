package service

import (
	"net/http"
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/entity"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/repository"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"

)

type PaymentService struct {
	InvoiceRepository   repository.InvoiceRepository
	CommunityRepository repository.CommunityRepository
	Supabase            *supabasestorageuploader.Client
	Client              snap.Client
}

func NewPayment(ir repository.InvoiceRepository,
	cr repository.CommunityRepository,
	supahase *supabasestorageuploader.Client) PaymentService {
	var client snap.Client
	env := midtrans.Sandbox
	client.New(os.Getenv("MIDTRANS_KEY"), env)

	return PaymentService{
		InvoiceRepository:   ir,
		Client:              client,
		CommunityRepository: cr,
		Supabase:            supahase,
	}
}

func (s *PaymentService) GenerateUrlToken(user model.UserTokenData, req model.CommunityRequest) (model.ServiceResponse, error) {
	community := entity.Community{
		ID:          uuid.NewString(),
		Name:        req.Name,
		UserID:      user.ID,
		Description: req.Description,
		Status:      "pending",
	}

	req.ProfilePicture.Filename = "community-pp" + community.ID + ".png"
	req.CoverPicture.Filename = "community-cp" + community.ID + ".png"

	ppUrl, err := s.Supabase.Upload(req.ProfilePicture)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to upload profile picture",
		}, err
	}
	cpUrl, err := s.Supabase.Upload(req.CoverPicture)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to upload cover picture",
		}, err
	}

	invoice := entity.Invoice{
		ID:      uuid.NewString(),
		UserID:  user.ID,
		Amount:  35000,
		Purpose: "community",
		Status:  "pending",
	}

	if err := s.InvoiceRepository.Insert(invoice); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to save invoice to database",
		}, err
	}

	community.ProfilePicture = ppUrl
	community.CoverPicture = cpUrl
	community.InvoiceID = invoice.ID

	if err := s.CommunityRepository.Insert(community); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to save community to database",
		}, err
	}

	payReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  invoice.ID,
			GrossAmt: int64(invoice.Amount),
		},
		Expiry: &snap.ExpiryDetails{
			Duration: 15,
			Unit:     "minute",
		},
	}

	snapResp, _ := s.Client.CreateTransaction(payReq)

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully create payment",
		Payload: gin.H{
			"snap_url": snapResp.RedirectURL,
		},
	}, nil
}

func (s *PaymentService) VerifyPayment(orderId string) bool {
	var client coreapi.Client
	env := midtrans.Sandbox
	client.New(os.Getenv("MIDTRANS_KEY"), env)

	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					return false
				} else if transactionStatusResp.FraudStatus == "accept" {
					return true
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return true
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				return false
			}
		}
	}
	return false
}
