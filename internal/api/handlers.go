package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PlayPixel/api/pkg/storage/mailbox"
)

func (s *APIServer) Mailbox(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		body := new(mailbox.Mailbox)
		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			return err
		}

		if body.Payload == "" {
			return fmt.Errorf("missing payload")
		}

		switch body.Payload {
		case "INSERT_MAILBOX":
			if err := mailbox.AddMailbox(body, s.ctx, s.store, s.log); err != nil {
				return err
			}

			return s.WriteJSON(w, http.StatusOK, ApiResponse{
				Success: true,
				Data:    "Mailbox Inserted",
			})
		case "READ_MAILBOX":
			mailboxData, err := mailbox.ReadMailbox(body, s.ctx, s.store, s.log)
			if err != nil {
				return err
			}

			return s.WriteJSON(w, http.StatusOK, ApiResponse{
				Success: true,
				Data:    mailboxData,
			})
		case "CLAIM_MAILBOX":
			if err := mailbox.DeleteMailbox(body, s.ctx, s.store, s.log); err != nil {
				return err
			}

			return s.WriteJSON(w, http.StatusOK, ApiResponse{
				Success: true,
				Data:    "Mailbox Deleted",
			})
		default:
			return fmt.Errorf("invalid payload")
		}
	}

	return fmt.Errorf("invalid method")
}

func (s *APIServer) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	return s.WriteJSON(w, http.StatusOK, ApiResponse{
		Success: true,
		Data:    "OK",
	})
}

// func InsertPlayer(w http.ResponseWriter, r *http.Request, s *APIServer) error {
// 	createAccReq := new(models.Account)
// 	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
// 		return err
// 	}

// 	if createAccReq.ID == 0 && createAccReq.Name == "" {
// 		return fmt.Errorf("ID cannot be 0")
// 	}

// 	account := models.NewPlayer(
// 		createAccReq.ID,
// 		createAccReq.Name,
// 		createAccReq.Secrets,
// 		createAccReq.Eggs,
// 		createAccReq.Bubbles,
// 		createAccReq.Power,
// 		createAccReq.Robux,
// 		createAccReq.Playtime,
// 	)

// 	if err := s.store.InsertAccounts(account); err != nil {
// 		return err
// 	}

// 	return s.WriteJSON(w, 200, ApiResponse{
// 		Success: true,
// 		Data:    account,
// 	})
// }

// func LeaderboardLookup(w http.ResponseWriter, r *http.Request, s *APIServer) error {
// 	if r.Method == "POST" {
// 		InsertAcc := new(models.AccountLookup)

// 		if err := json.NewDecoder(r.Body).Decode(InsertAcc); err != nil {
// 			return err
// 		}

// 		if InsertAcc.RobloxID == 0 {
// 			return fmt.Errorf("Missing robloxId")
// 		}

// 		acc, err := s.store.GetSpecificPlayer(InsertAcc.RobloxID)

// 		if err != nil {
// 			return err
// 		}

// 		return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 			Success: true,
// 			Data:    acc,
// 		})
// 	}

// 	return fmt.Errorf("Invalid Method")
// }

// func PetsExistance(w http.ResponseWriter, r *http.Request, s *APIServer) error {
// 	if r.Method == "POST" {
// 		InsertAcc := new(models.PetsExistance)

// 		if err := json.NewDecoder(r.Body).Decode(InsertAcc); err != nil {
// 			return err
// 		}

// 		if InsertAcc.Payload == "" {
// 			return fmt.Errorf("Missing Payload")
// 		}

// 		switch InsertAcc.Payload {
// 		case "INSERT_PETS_EXISTANCE":
// 			if err := s.store.InsertPetsExistance(InsertAcc); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Pets Inserted",
// 			})
// 		case "READ_PETS_EXISTANCE":
// 			pets, err := s.store.GetPetsExistance()

// 			if err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    pets,
// 			})
// 		case "DELETE_PETS_EXISTANCE":
// 			if err := s.store.DeletePetsExistence(InsertAcc); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Pets Removed",
// 			})
// 		default:
// 			return fmt.Errorf("Invalid Payload")
// 		}
// 	}

// 	return fmt.Errorf("Invalid Method")
// }

// func Auctions(w http.ResponseWriter, r *http.Request, s *APIServer) error {
// 	if r.Method == "POST" {
// 		Auction := new(models.AuctionAccount)
// 		if err := json.NewDecoder(r.Body).Decode(Auction); err != nil {
// 			return err
// 		}

// 		if Auction.Payload == "" {
// 			return fmt.Errorf("Invalid Payload")
// 		}

// 		switch Auction.Payload {
// 		case "LIST":
// 			if err := s.store.ListAuction(Auction); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Auction Inserted",
// 			})

// 		case "READ":
// 			auctions, err := s.store.GetAuctions()
// 			if err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    auctions,
// 			})
// 		case "DELETE":
// 			if err := s.store.RemoveAuction(Auction); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Auction Deleted",
// 			})
// 		case "PURCHASE":
// 			if err := s.store.PurchaseAuction(Auction); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Auction Purchased",
// 			})
// 		case "AUCTION_GET_CLAIMS":
// 			claims, err := s.store.GetAuctionClaims(Auction)

// 			if err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    claims,
// 			})
// 		case "AUCTION_CLAIM":
// 			if err := s.store.AuctionClaim(Auction); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Auction Claimed",
// 			})
// 		case "AUCTION_GET_LISTINGS":
// 			listing, err := s.store.GetAuctionListing(Auction)

// 			if err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    listing,
// 			})
// 		case "AUCTION_UNLIST":
// 			if err := s.store.AuctionUnlist(Auction); err != nil {
// 				return err
// 			}

// 			return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 				Success: true,
// 				Data:    "Auction Unlisted",
// 			})
// 		}
// 	}

// 	return fmt.Errorf("Invalid Method")
// }

// func GetLeaderboards(w http.ResponseWriter, r *http.Request, s *APIServer) error {
// 	vars := mux.Vars(r)["which"]

// 	leaderboards := map[string]func() (*models.PlayerDataResponse, error){
// 		"eggs":     s.store.GetEggs,
// 		"bubbles":  s.store.GetBubbles,
// 		"secrets":  s.store.GetSecrets,
// 		"power":    s.store.GetPower,
// 		"robux":    s.store.GetRobux,
// 		"playtime": s.store.GetPlaytime,
// 	}

// 	if leaderboardFunc, ok := leaderboards[vars]; ok {
// 		data, err := leaderboardFunc()
// 		if err != nil {
// 			return err
// 		}

// 		return s.WriteJSON(w, http.StatusOK, ApiResponse{
// 			Success: true,
// 			Data:    data,
// 		})
// 	}

// 	return fmt.Errorf("Invalid Leaderboard")
// }
