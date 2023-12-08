package devices

import (
	"encoding/json"
	"net/http"
	_"time"

	"github.com/darthsalad/univboard/internal/logger"
	"github.com/darthsalad/univboard/internal/utils"
	"github.com/darthsalad/univboard/pkg/database"
	"github.com/darthsalad/univboard/pkg/models"
)

//TODO routes(devices)
//POST /devices - Done
//GET /devices - Done
//GET /devices/{id}
//PUT /devices/{id}
//DELETE /devices/{id}


func AddDevice(db *database.Database, w http.ResponseWriter, r *http.Request) error {
	device := models.Device{}
	userId := r.Context().Value("user_id").(string)

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		logger.Fatalf("err decoding: %v", err)
		return err
	}

	device = *models.NewDevice(
		userId,
		device.Name,
		device.OS,
		device.OSVersion,
	)

	if err := db.AddDevice(&device); err != nil {
		logger.Logf("err adding device: %v", err)
		utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
			"error": map[string]any{
				"message":     err.Error(),
				"status_code": http.StatusInternalServerError,
			},
		})
		return nil
	}

	err := utils.JsonResp(w, http.StatusCreated, map[string]any{
		"message": "Successfully added device",
		"device": map[string]string{
			"name":       device.Name,
			"os":         device.OS,
			"os_version": device.OSVersion,
			"created_at": device.CreatedAt,
		},
	})
	if err != nil {
		logger.Logf("err responding: %v", err)
		return err
	}

	return nil
}

func GetDeviceByID(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := r.Context().Value("id").(string)

    device, err := db.GetDeviceByID(id)
    if err != nil {
        logger.Logf("error getting device: %v", err)
        utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
            "error": map[string]any{
                "message":     err.Error(),
                "status_code": http.StatusInternalServerError,
            },
        })
        return err
    }

    err = utils.JsonResp(w, http.StatusOK, device)
    if err != nil {
        logger.Logf("err responding: %v", err)
        return err
    }

    return nil
}

func UpdateDevice(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := r.Context().Value("id").(string)
    device := models.Device{}

    if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
        logger.Fatalf("err decoding: %v", err)
        return err
    }

    if err := db.UpdateDevice(id, &device); err != nil {
        logger.Logf("err updating device: %v", err)
        utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
            "error": map[string]any{
                "message":     err.Error(),
                "status_code": http.StatusInternalServerError,
            },
        })
        return err
    }

    err := utils.JsonResp(w, http.StatusOK, device)
    if err != nil {
        logger.Logf("err responding: %v", err)
        return err
    }

    return nil
}

func RemoveDevice(db *database.Database, w http.ResponseWriter, r *http.Request) error {
    id := r.Context().Value("id").(string)

    if err := db.RemoveDevice(id); err != nil {
        logger.Logf("err removing device: %v", err)
        utils.JsonResp(w, http.StatusInternalServerError, map[string]any{
            "error": map[string]any{
                "message":     err.Error(),
                "status_code": http.StatusInternalServerError,
            },
        })
        return err
    }

    err := utils.JsonResp(w, http.StatusOK, map[string]string{
        "message": "Successfully removed device",
    })
    if err != nil {
        logger.Logf("err responding: %v", err)
        return err
    }

    return nil
}