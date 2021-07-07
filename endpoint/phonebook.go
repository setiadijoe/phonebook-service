package endpoint

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetList ...
func MakeGetList(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetListRequest)
		params := &model.ParamsPhoneBook{
			Status:     req.Status,
			Search:     helper.SetPointerString(req.Search),
			RegencyID:  helper.SetPointerInt64(req.RegencyID),
			DistrictID: helper.SetPointerInt64(req.DistrictID),
			VillageID:  helper.SetPointerInt64(req.VillageID),
			Limit:      helper.SetPointerInt64(req.Limit),
			Page:       helper.SetPointerInt64(req.Page),
			Longitude:  helper.SetPointerString(req.Longitude),
			Latitude:   helper.SetPointerString(req.Latitude),
			SortBy:     helper.SetPointerString(req.SortBy),
			OrderBy:    helper.SetPointerString(helper.AscOrDesc[req.OrderBy]),
			Name:       helper.SetPointerString(req.Name),
			Phone:      helper.SetPointerString(req.PhoneNumber),
		}

		resp, err := usecase.GetList(ctx, params)

		if err != nil {
			return nil, err
		}

		phonebooks := EncodePhonebook(resp.PhoneBooks)

		return &PhoneBookWithMeta{
			Data: &PhonebookWithMeta{
				Phonebooks: phonebooks,
				Meta: &Metadata{
					TotalCount:  resp.Metadata.TotalCount,
					PageCount:   resp.Metadata.PageCount,
					CurrentPage: resp.Metadata.CurrentPage,
					PerPage:     resp.Metadata.PerPage,
				},
			},
		}, nil
	}
}

// MakeGetDetail ...
func MakeGetDetail(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetDetailRequest)
		resp, err := usecase.GetDetail(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		phoneNumbers := []*PhoneNumber{}
		if err := json.Unmarshal([]byte(resp.PhoneNumbers), &phoneNumbers); err != nil {
			return nil, err
		}

		data := &PhonebookDetail{
			ID:             resp.ID,
			Name:           resp.Name,
			Category:       resp.Category,
			Address:        resp.Address,
			Description:    resp.Description,
			PhoneNumbers:   phoneNumbers,
			Regency:        resp.Regency,
			District:       resp.District,
			Village:        resp.Village,
			Latitude:       resp.Latitude,
			Longitude:      resp.Longitude,
			CoverImagePath: fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, resp.CoverImageURL),
			Sequence:       resp.Sequence,
			Status:         resp.Status,
			StatusLabel:    GetStatusLabel[resp.Status]["id"],
			CreatedAt:      resp.CreatedAt,
			UpdatedAt:      resp.UpdatedAt,
		}
		return map[string]interface{}{
			"data": data,
		}, nil
	}
}

// MakeAddPhonebook ...
func MakeAddPhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*AddPhonebookRequest)
		phoneNumbers, _ := json.Marshal(req.PhoneNumbers)
		if err := usecase.Insert(ctx, &model.AddPhonebook{
			Name:           req.Name,
			PhoneNumbers:   helper.SetPointerString(string(phoneNumbers)),
			Address:        req.Address,
			Description:    req.Description,
			RegencyID:      req.RegencyID,
			DistrictID:     req.DistrictID,
			VillageID:      req.VillageID,
			CategoryID:     req.CategoryID,
			Latitude:       req.Latitude,
			Longitude:      req.Longitude,
			CoverImagePath: req.CoverImagePath,
			Status:         req.Status,
			Sequence:       req.Sequence,
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "phonebook_has_been_created_succesfully",
		}, nil

	}
}

// MakeUpdatePhonebook ...
func MakeUpdatePhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdatePhonebookRequest)
		phoneNumbers, _ := json.Marshal(req.PhoneNumbers)
		if err := usecase.Update(ctx, &model.UpdatePhonebook{
			ID:             req.ID,
			Name:           req.Name,
			PhoneNumbers:   helper.SetPointerString(string(phoneNumbers)),
			Address:        req.Address,
			Description:    req.Description,
			RegencyID:      req.RegencyID,
			DistrictID:     req.DistrictID,
			VillageID:      req.VillageID,
			CategoryID:     req.CategoryID,
			Latitude:       req.Latitude,
			Longitude:      req.Longitude,
			CoverImagePath: req.CoverImagePath,
			Status:         req.Status,
			Sequence:       req.Sequence,
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_UPDATED,
			Message: "phonebook_has_been_updated_successfuly",
		}, nil
	}
}

// MakeDeletePhonebook ...
func MakeDeletePhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetDetailRequest)
		if err := usecase.Delete(ctx, req.ID); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_DELETED,
			Message: "phonebook_has_been_deleted_successfuly",
		}, nil
	}
}

func MakeCheckHealthy(ctx context.Context) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ok",
		}, nil
	}
}

func MakeCheckReadiness(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if err := usecase.CheckHealthReadiness(ctx); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ready",
		}, nil
	}
}

// MakeIsExistPhoneNumber ...
func MakeIsExistPhoneNumber(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*IsExistPhoneNumber)

		isExist, err := usecase.IsExistPhoneNumber(ctx, req.PhoneNumber)
		if err != nil {
			return nil, err
		}

		return map[string]map[string]interface{}{
			"data": {"exist": isExist},
		}, nil
	}
}
