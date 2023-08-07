package model

import "time"

type LoginByEmailRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	DeviceID   string `json:"device_id"`
	RememberMe bool   `json:"remember_me"`
}

type LoginByEmailResponse struct {
	Status     int `json:"status"`
	LoginBasic struct {
		Tokens Token `json:"tokens"`
	} `json:"login_basic"`
}

type LoginByAPIKeyRequest struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type LoginByAPIKeyResponse struct {
	Status       int `json:"status"`
	ApiKeysLogin struct {
		Tokens Token `json:"tokens"`
	} `json:"api_keys_login"`
}

type CreateAPIKeyRequest struct {
	AcSameAsForUser            bool   `json:"ac_same_as_for_user"`
	Branches                   []int  `json:"branches"`
	Firms                      []int  `json:"firms"`
	HasAccessToAllFirmBranches bool   `json:"has_access_to_all_firm_branches"`
	Name                       string `json:"name"`
	Permissions                []int  `json:"permissions"`
	Roles                      []int  `json:"roles"`
	UpdatedAt                  string `json:"updated_at"`
}

type GetB2BUsersV2Request struct {
	IncludeChiefadmin bool                            `json:"include_chiefadmin"`
	Filters           GetB2BUserByIDFilters           `json:"filters"`
	Pagination        GetB2BUserByIDFiltersPagination `json:"pagination,omitempty"`
	Search            string                          `json:"search,omitempty"`
	Sorts             []GetB2BUserByIDFiltersSort     `json:"sorts,omitempty"`
}

type GetB2BUserByIDFilters struct {
	BranchIds       []int  `json:"branch_ids,omitempty"`
	EmailVerified   bool   `json:"email_verified,omitempty"`
	FirmIds         []int  `json:"firm_ids,omitempty"`
	Ids             []int  `json:"ids,omitempty"`
	MicroserviceIds []int  `json:"microservice_ids,omitempty"`
	NeitherVerified bool   `json:"neither_verified"`
	PhoneVerified   bool   `json:"phone_verified"`
	RoleIds         []int  `json:"role_ids,omitempty"`
	ShowOnly        string `json:"show_only,omitempty"`
}

type GetB2BUserByIDFiltersPagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type GetB2BUserByIDFiltersSort struct {
	Direction string `json:"direction"`
	Field     string `json:"field"`
}

type GetB2BUsersV2Response struct {
	CurrentPage  int `json:"current_page"`
	PagesCount   int `json:"pages_count"`
	Status       int `json:"status"`
	TotalRecords int `json:"total_records"`
	UsersList    []struct {
		AllowFeedback bool `json:"allow_feedback"`
		Branches      []struct {
			Code   string `json:"code"`
			FirmID int    `json:"firm_id"`
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Slug   string `json:"slug"`
		} `json:"branches"`
		CreatedAt        time.Time `json:"created_at"`
		DoNotAskMfaAgain bool      `json:"do_not_ask_mfa_again"`
		Email            string    `json:"email"`
		EmailVerified    bool      `json:"email_verified"`
		Firms            []struct {
			AdminsCount                int  `json:"admins_count"`
			AllowB2CEmailNotifications bool `json:"allow_b2c_email_notifications"`
			Branches                   []struct {
				Code   string `json:"code"`
				FirmID int    `json:"firm_id"`
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Slug   string `json:"slug"`
			} `json:"branches"`
			BranchesCount     int       `json:"branches_count"`
			CreatedAt         time.Time `json:"created_at"`
			DefaultCompliance struct {
				AllowFeedback              bool      `json:"allow_feedback"`
				BigAvatarURL               string    `json:"big_avatar_url"`
				CreatedAt                  time.Time `json:"created_at"`
				DoNotAskMfaAgain           bool      `json:"do_not_ask_mfa_again"`
				Email                      string    `json:"email"`
				EmailVerified              bool      `json:"email_verified"`
				FirstName                  string    `json:"first_name"`
				HasAccessToAllFirmBranches bool      `json:"has_access_to_all_firm_branches"`
				ID                         int       `json:"id"`
				InitiatorID                int       `json:"initiator_id"`
				IsChiefAdmin               bool      `json:"is_chief_admin"`
				LastName                   string    `json:"last_name"`
				Locale                     struct {
					ID     int    `json:"id"`
					Locale string `json:"locale"`
					Name   string `json:"name"`
				} `json:"locale"`
				LocaleID      int  `json:"locale_id"`
				MfaEnabled    bool `json:"mfa_enabled"`
				Microservices []struct {
					Entities struct {
						AdditionalProp1 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp1"`
						AdditionalProp2 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp2"`
						AdditionalProp3 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp3"`
					} `json:"entities"`
					ID                      int    `json:"id"`
					IsActive                bool   `json:"is_active"`
					IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
					LogoURL                 string `json:"logo_url"`
					Name                    string `json:"name"`
					Slug                    string `json:"slug"`
				} `json:"microservices"`
				Mode           string    `json:"mode"`
				Phone          string    `json:"phone"`
				PhoneVerified  bool      `json:"phone_verified"`
				RepCode        string    `json:"rep_code"`
				SmallAvatarURL string    `json:"small_avatar_url"`
				Theme          string    `json:"theme"`
				UpdatedAt      time.Time `json:"updated_at"`
			} `json:"default_compliance"`
			DefaultComplianceID int       `json:"default_compliance_id"`
			ID                  int       `json:"id"`
			LastStatusChangeAt  string    `json:"last_status_change_at"`
			Name                string    `json:"name"`
			Slug                string    `json:"slug"`
			Status              string    `json:"status"`
			UpdatedAt           time.Time `json:"updated_at"`
		} `json:"firms"`
		FirstName   string `json:"first_name"`
		ID          int    `json:"id"`
		InitiatorID int    `json:"initiator_id"`
		LastName    string `json:"last_name"`
		Locale      struct {
			ID     int    `json:"id"`
			Locale string `json:"locale"`
			Name   string `json:"name"`
		} `json:"locale"`
		LocaleID      int  `json:"locale_id"`
		MfaEnabled    bool `json:"mfa_enabled"`
		Microservices []struct {
			CreatedAt string `json:"created_at"`
			Domains   []struct {
				InternalDomain   string `json:"internal_domain"`
				InternalProtocol string `json:"internal_protocol"`
				PublicPathPrefix string `json:"public_path_prefix"`
			} `json:"domains"`
			ID                      int    `json:"id"`
			IsActive                bool   `json:"is_active"`
			IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
			LogoURL                 string `json:"logo_url"`
			Name                    string `json:"name"`
			Slug                    string `json:"slug"`
			UpdatedAt               string `json:"updated_at"`
		} `json:"microservices"`
		Permissions []struct {
			AssignmentGroup string `json:"assignment_group"`
			CreatedAt       string `json:"created_at"`
			ID              int    `json:"id"`
			Metadata        struct {
				Entities struct {
				} `json:"entities"`
				IsMicroserviceGlobal bool `json:"is_microservice_global"`
			} `json:"metadata"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				ID                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoURL                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceID int    `json:"microservice_id"`
			Permission     string `json:"permission"`
			UpdatedAt      string `json:"updated_at"`
		} `json:"permissions"`
		Phone         string `json:"phone"`
		PhoneVerified bool   `json:"phone_verified"`
		RepCode       string `json:"rep_code"`
		Roles         []struct {
			CreatedAt string `json:"created_at"`
			Firm      struct {
				AdminsCount                int  `json:"admins_count"`
				AllowB2CEmailNotifications bool `json:"allow_b2c_email_notifications"`
				Branches                   []struct {
					Code   string `json:"code"`
					FirmID int    `json:"firm_id"`
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Slug   string `json:"slug"`
				} `json:"branches"`
				BranchesCount     int       `json:"branches_count"`
				CreatedAt         time.Time `json:"created_at"`
				DefaultCompliance struct {
					AllowFeedback              bool      `json:"allow_feedback"`
					BigAvatarURL               string    `json:"big_avatar_url"`
					CreatedAt                  time.Time `json:"created_at"`
					DoNotAskMfaAgain           bool      `json:"do_not_ask_mfa_again"`
					Email                      string    `json:"email"`
					EmailVerified              bool      `json:"email_verified"`
					FirstName                  string    `json:"first_name"`
					HasAccessToAllFirmBranches bool      `json:"has_access_to_all_firm_branches"`
					ID                         int       `json:"id"`
					InitiatorID                int       `json:"initiator_id"`
					IsChiefAdmin               bool      `json:"is_chief_admin"`
					LastName                   string    `json:"last_name"`
					Locale                     struct {
						ID     int    `json:"id"`
						Locale string `json:"locale"`
						Name   string `json:"name"`
					} `json:"locale"`
					LocaleID      int  `json:"locale_id"`
					MfaEnabled    bool `json:"mfa_enabled"`
					Microservices []struct {
						Entities struct {
							AdditionalProp1 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp1"`
							AdditionalProp2 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp2"`
							AdditionalProp3 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp3"`
						} `json:"entities"`
						ID                      int    `json:"id"`
						IsActive                bool   `json:"is_active"`
						IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
						LogoURL                 string `json:"logo_url"`
						Name                    string `json:"name"`
						Slug                    string `json:"slug"`
					} `json:"microservices"`
					Mode           string    `json:"mode"`
					Phone          string    `json:"phone"`
					PhoneVerified  bool      `json:"phone_verified"`
					RepCode        string    `json:"rep_code"`
					SmallAvatarURL string    `json:"small_avatar_url"`
					Theme          string    `json:"theme"`
					UpdatedAt      time.Time `json:"updated_at"`
				} `json:"default_compliance"`
				DefaultComplianceID int       `json:"default_compliance_id"`
				ID                  int       `json:"id"`
				LastStatusChangeAt  string    `json:"last_status_change_at"`
				Name                string    `json:"name"`
				Slug                string    `json:"slug"`
				Status              string    `json:"status"`
				UpdatedAt           time.Time `json:"updated_at"`
			} `json:"firm"`
			FirmID       int `json:"firm_id"`
			ID           int `json:"id"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				ID                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoURL                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceID int    `json:"microservice_id"`
			Name           string `json:"name"`
			Permissions    []struct {
				AssignmentGroup string `json:"assignment_group"`
				CreatedAt       string `json:"created_at"`
				ID              int    `json:"id"`
				Metadata        struct {
					Entities struct {
					} `json:"entities"`
					IsMicroserviceGlobal bool `json:"is_microservice_global"`
				} `json:"metadata"`
				Microservice struct {
					CreatedAt string `json:"created_at"`
					Domains   []struct {
						InternalDomain   string `json:"internal_domain"`
						InternalProtocol string `json:"internal_protocol"`
						PublicPathPrefix string `json:"public_path_prefix"`
					} `json:"domains"`
					ID                      int    `json:"id"`
					IsActive                bool   `json:"is_active"`
					IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
					LogoURL                 string `json:"logo_url"`
					Name                    string `json:"name"`
					Slug                    string `json:"slug"`
					UpdatedAt               string `json:"updated_at"`
				} `json:"microservice"`
				MicroserviceID int    `json:"microservice_id"`
				Permission     string `json:"permission"`
				UpdatedAt      string `json:"updated_at"`
			} `json:"permissions"`
			Type struct {
				Color     string `json:"color"`
				CreatedAt string `json:"created_at"`
				ID        int    `json:"id"`
				Type      string `json:"type"`
				UpdatedAt string `json:"updated_at"`
			} `json:"type"`
			TypeID    int    `json:"type_id"`
			UpdatedAt string `json:"updated_at"`
		} `json:"roles"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"users_list"`
}

type CreateAPIKeyResponse struct {
	ApiKeys struct {
		AcSameAsForUser bool   `json:"ac_same_as_for_user"`
		ApiKey          string `json:"api_key"`
		ApiSecret       string `json:"api_secret"`
		Branches        []struct {
			Code            string `json:"code"`
			FirmId          int    `json:"firm_id"`
			Id              int    `json:"id"`
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			TradingBranchId int    `json:"trading_branch_id"`
		} `json:"branches"`
		CreatedAt   string `json:"created_at"`
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Permissions []struct {
			AssignmentGroup string `json:"assignment_group"`
			CreatedAt       string `json:"created_at"`
			Id              int    `json:"id"`
			Metadata        struct {
				Entities struct {
				} `json:"entities"`
				IsMicroserviceGlobal bool `json:"is_microservice_global"`
			} `json:"metadata"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				Id                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoUrl                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceId int    `json:"microservice_id"`
			Permission     string `json:"permission"`
			UpdatedAt      string `json:"updated_at"`
		} `json:"permissions"`
		Roles []struct {
			CreatedAt string `json:"created_at"`
			Firm      struct {
				AdminsCount                int  `json:"admins_count"`
				AllowB2CEmailNotifications bool `json:"allow_b2c_email_notifications"`
				Branches                   []struct {
					Code            string `json:"code"`
					FirmId          int    `json:"firm_id"`
					Id              int    `json:"id"`
					Name            string `json:"name"`
					Slug            string `json:"slug"`
					TradingBranchId int    `json:"trading_branch_id"`
				} `json:"branches"`
				BranchesCount     int    `json:"branches_count"`
				CreatedAt         string `json:"created_at"`
				DefaultCompliance struct {
					AllowFeedback              bool   `json:"allow_feedback"`
					BigAvatarUrl               string `json:"big_avatar_url"`
					CreatedAt                  string `json:"created_at"`
					DoNotAskMfaAgain           bool   `json:"do_not_ask_mfa_again"`
					Email                      string `json:"email"`
					EmailVerified              bool   `json:"email_verified"`
					FirstName                  string `json:"first_name"`
					HasAccessToAllFirmBranches bool   `json:"has_access_to_all_firm_branches"`
					Id                         int    `json:"id"`
					InitiatorId                int    `json:"initiator_id"`
					IsChiefAdmin               bool   `json:"is_chief_admin"`
					LastName                   string `json:"last_name"`
					Locale                     struct {
						Id     int    `json:"id"`
						Locale string `json:"locale"`
						Name   string `json:"name"`
					} `json:"locale"`
					LocaleId      int  `json:"locale_id"`
					MfaEnabled    bool `json:"mfa_enabled"`
					Microservices []struct {
						Entities struct {
							AdditionalProp1 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp1"`
							AdditionalProp2 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp2"`
							AdditionalProp3 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp3"`
						} `json:"entities"`
						Id                      int    `json:"id"`
						IsActive                bool   `json:"is_active"`
						IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
						LogoUrl                 string `json:"logo_url"`
						Name                    string `json:"name"`
						Slug                    string `json:"slug"`
					} `json:"microservices"`
					Mode           string `json:"mode"`
					Phone          string `json:"phone"`
					PhoneVerified  bool   `json:"phone_verified"`
					RepCode        string `json:"rep_code"`
					SmallAvatarUrl string `json:"small_avatar_url"`
					Theme          string `json:"theme"`
					UpdatedAt      string `json:"updated_at"`
				} `json:"default_compliance"`
				DefaultComplianceId int    `json:"default_compliance_id"`
				Id                  int    `json:"id"`
				LastStatusChangeAt  string `json:"last_status_change_at"`
				Name                string `json:"name"`
				Slug                string `json:"slug"`
				Status              string `json:"status"`
				UpdatedAt           string `json:"updated_at"`
			} `json:"firm"`
			FirmId       int `json:"firm_id"`
			Id           int `json:"id"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				Id                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoUrl                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceId int    `json:"microservice_id"`
			Name           string `json:"name"`
			Permissions    []struct {
				AssignmentGroup string `json:"assignment_group"`
				CreatedAt       string `json:"created_at"`
				Id              int    `json:"id"`
				Metadata        struct {
					Entities struct {
					} `json:"entities"`
					IsMicroserviceGlobal bool `json:"is_microservice_global"`
				} `json:"metadata"`
				Microservice struct {
					CreatedAt string `json:"created_at"`
					Domains   []struct {
						InternalDomain   string `json:"internal_domain"`
						InternalProtocol string `json:"internal_protocol"`
						PublicPathPrefix string `json:"public_path_prefix"`
					} `json:"domains"`
					Id                      int    `json:"id"`
					IsActive                bool   `json:"is_active"`
					IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
					LogoUrl                 string `json:"logo_url"`
					Name                    string `json:"name"`
					Slug                    string `json:"slug"`
					UpdatedAt               string `json:"updated_at"`
				} `json:"microservice"`
				MicroserviceId int    `json:"microservice_id"`
				Permission     string `json:"permission"`
				UpdatedAt      string `json:"updated_at"`
			} `json:"permissions"`
			Type struct {
				Color     string `json:"color"`
				CreatedAt string `json:"created_at"`
				Id        int    `json:"id"`
				Type      string `json:"type"`
				UpdatedAt string `json:"updated_at"`
			} `json:"type"`
			TypeId    int    `json:"type_id"`
			UpdatedAt string `json:"updated_at"`
		} `json:"roles"`
		UpdatedAt string `json:"updated_at"`
	} `json:"api_keys"`
	Status int `json:"status"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type GetB2BUserByIDResponse struct {
	B2BUser struct {
		AllowFeedback bool   `json:"allow_feedback"`
		BigAvatarUrl  string `json:"big_avatar_url"`
		Branches      []struct {
			Code            string `json:"code"`
			FirmId          int    `json:"firm_id"`
			Id              int    `json:"id"`
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			TradingBranchId int    `json:"trading_branch_id"`
		} `json:"branches"`
		CreatedAt        string `json:"created_at"`
		DoNotAskMfaAgain bool   `json:"do_not_ask_mfa_again"`
		Email            string `json:"email"`
		EmailVerified    bool   `json:"email_verified"`
		Firms            []struct {
			AdminsCount                int  `json:"admins_count"`
			AllowB2CEmailNotifications bool `json:"allow_b2c_email_notifications"`
			Branches                   []struct {
				Code            string `json:"code"`
				FirmId          int    `json:"firm_id"`
				Id              int    `json:"id"`
				Name            string `json:"name"`
				Slug            string `json:"slug"`
				TradingBranchId int    `json:"trading_branch_id"`
			} `json:"branches"`
			BranchesCount     int    `json:"branches_count"`
			CreatedAt         string `json:"created_at"`
			DefaultCompliance struct {
				AllowFeedback              bool   `json:"allow_feedback"`
				BigAvatarUrl               string `json:"big_avatar_url"`
				CreatedAt                  string `json:"created_at"`
				DoNotAskMfaAgain           bool   `json:"do_not_ask_mfa_again"`
				Email                      string `json:"email"`
				EmailVerified              bool   `json:"email_verified"`
				FirstName                  string `json:"first_name"`
				HasAccessToAllFirmBranches bool   `json:"has_access_to_all_firm_branches"`
				Id                         int    `json:"id"`
				InitiatorId                int    `json:"initiator_id"`
				IsChiefAdmin               bool   `json:"is_chief_admin"`
				LastName                   string `json:"last_name"`
				Locale                     struct {
					Id     int    `json:"id"`
					Locale string `json:"locale"`
					Name   string `json:"name"`
				} `json:"locale"`
				LocaleId      int  `json:"locale_id"`
				MfaEnabled    bool `json:"mfa_enabled"`
				Microservices []struct {
					Entities struct {
						AdditionalProp1 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp1"`
						AdditionalProp2 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp2"`
						AdditionalProp3 struct {
							AdditionalProp1 bool `json:"additionalProp1"`
							AdditionalProp2 bool `json:"additionalProp2"`
							AdditionalProp3 bool `json:"additionalProp3"`
						} `json:"additionalProp3"`
					} `json:"entities"`
					Id                      int    `json:"id"`
					IsActive                bool   `json:"is_active"`
					IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
					LogoUrl                 string `json:"logo_url"`
					Name                    string `json:"name"`
					Slug                    string `json:"slug"`
				} `json:"microservices"`
				Mode           string `json:"mode"`
				Phone          string `json:"phone"`
				PhoneVerified  bool   `json:"phone_verified"`
				RepCode        string `json:"rep_code"`
				SmallAvatarUrl string `json:"small_avatar_url"`
				Theme          string `json:"theme"`
				UpdatedAt      string `json:"updated_at"`
			} `json:"default_compliance"`
			DefaultComplianceId int    `json:"default_compliance_id"`
			Id                  int    `json:"id"`
			LastStatusChangeAt  string `json:"last_status_change_at"`
			Name                string `json:"name"`
			Slug                string `json:"slug"`
			Status              string `json:"status"`
			UpdatedAt           string `json:"updated_at"`
		} `json:"firms"`
		FirstName                  string `json:"first_name"`
		HasAccessToAllFirmBranches bool   `json:"has_access_to_all_firm_branches"`
		Id                         int    `json:"id"`
		InitiatorId                int    `json:"initiator_id"`
		LastName                   string `json:"last_name"`
		Locale                     struct {
			Id     int    `json:"id"`
			Locale string `json:"locale"`
			Name   string `json:"name"`
		} `json:"locale"`
		LocaleId      int  `json:"locale_id"`
		MfaEnabled    bool `json:"mfa_enabled"`
		Microservices []struct {
			CreatedAt string `json:"created_at"`
			Domains   []struct {
				InternalDomain   string `json:"internal_domain"`
				InternalProtocol string `json:"internal_protocol"`
				PublicPathPrefix string `json:"public_path_prefix"`
			} `json:"domains"`
			Id                      int    `json:"id"`
			IsActive                bool   `json:"is_active"`
			IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
			LogoUrl                 string `json:"logo_url"`
			Name                    string `json:"name"`
			Slug                    string `json:"slug"`
			UpdatedAt               string `json:"updated_at"`
		} `json:"microservices"`
		Mode        string `json:"mode"`
		Permissions []struct {
			AssignmentGroup string `json:"assignment_group"`
			CreatedAt       string `json:"created_at"`
			Id              int    `json:"id"`
			Metadata        struct {
				Entities struct {
				} `json:"entities"`
				IsMicroserviceGlobal bool `json:"is_microservice_global"`
			} `json:"metadata"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				Id                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoUrl                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceId int    `json:"microservice_id"`
			Permission     string `json:"permission"`
			UpdatedAt      string `json:"updated_at"`
		} `json:"permissions"`
		Phone         string `json:"phone"`
		PhoneVerified bool   `json:"phone_verified"`
		RepCode       string `json:"rep_code"`
		Roles         []struct {
			CreatedAt string `json:"created_at"`
			Firm      struct {
				AdminsCount                int  `json:"admins_count"`
				AllowB2CEmailNotifications bool `json:"allow_b2c_email_notifications"`
				Branches                   []struct {
					Code            string `json:"code"`
					FirmId          int    `json:"firm_id"`
					Id              int    `json:"id"`
					Name            string `json:"name"`
					Slug            string `json:"slug"`
					TradingBranchId int    `json:"trading_branch_id"`
				} `json:"branches"`
				BranchesCount     int    `json:"branches_count"`
				CreatedAt         string `json:"created_at"`
				DefaultCompliance struct {
					AllowFeedback              bool   `json:"allow_feedback"`
					BigAvatarUrl               string `json:"big_avatar_url"`
					CreatedAt                  string `json:"created_at"`
					DoNotAskMfaAgain           bool   `json:"do_not_ask_mfa_again"`
					Email                      string `json:"email"`
					EmailVerified              bool   `json:"email_verified"`
					FirstName                  string `json:"first_name"`
					HasAccessToAllFirmBranches bool   `json:"has_access_to_all_firm_branches"`
					Id                         int    `json:"id"`
					InitiatorId                int    `json:"initiator_id"`
					IsChiefAdmin               bool   `json:"is_chief_admin"`
					LastName                   string `json:"last_name"`
					Locale                     struct {
						Id     int    `json:"id"`
						Locale string `json:"locale"`
						Name   string `json:"name"`
					} `json:"locale"`
					LocaleId      int  `json:"locale_id"`
					MfaEnabled    bool `json:"mfa_enabled"`
					Microservices []struct {
						Entities struct {
							AdditionalProp1 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp1"`
							AdditionalProp2 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp2"`
							AdditionalProp3 struct {
								AdditionalProp1 bool `json:"additionalProp1"`
								AdditionalProp2 bool `json:"additionalProp2"`
								AdditionalProp3 bool `json:"additionalProp3"`
							} `json:"additionalProp3"`
						} `json:"entities"`
						Id                      int    `json:"id"`
						IsActive                bool   `json:"is_active"`
						IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
						LogoUrl                 string `json:"logo_url"`
						Name                    string `json:"name"`
						Slug                    string `json:"slug"`
					} `json:"microservices"`
					Mode           string `json:"mode"`
					Phone          string `json:"phone"`
					PhoneVerified  bool   `json:"phone_verified"`
					RepCode        string `json:"rep_code"`
					SmallAvatarUrl string `json:"small_avatar_url"`
					Theme          string `json:"theme"`
					UpdatedAt      string `json:"updated_at"`
				} `json:"default_compliance"`
				DefaultComplianceId int    `json:"default_compliance_id"`
				Id                  int    `json:"id"`
				LastStatusChangeAt  string `json:"last_status_change_at"`
				Name                string `json:"name"`
				Slug                string `json:"slug"`
				Status              string `json:"status"`
				UpdatedAt           string `json:"updated_at"`
			} `json:"firm"`
			FirmId       int `json:"firm_id"`
			Id           int `json:"id"`
			Microservice struct {
				CreatedAt string `json:"created_at"`
				Domains   []struct {
					InternalDomain   string `json:"internal_domain"`
					InternalProtocol string `json:"internal_protocol"`
					PublicPathPrefix string `json:"public_path_prefix"`
				} `json:"domains"`
				Id                      int    `json:"id"`
				IsActive                bool   `json:"is_active"`
				IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
				LogoUrl                 string `json:"logo_url"`
				Name                    string `json:"name"`
				Slug                    string `json:"slug"`
				UpdatedAt               string `json:"updated_at"`
			} `json:"microservice"`
			MicroserviceId int    `json:"microservice_id"`
			Name           string `json:"name"`
			Permissions    []struct {
				AssignmentGroup string `json:"assignment_group"`
				CreatedAt       string `json:"created_at"`
				Id              int    `json:"id"`
				Metadata        struct {
					Entities struct {
					} `json:"entities"`
					IsMicroserviceGlobal bool `json:"is_microservice_global"`
				} `json:"metadata"`
				Microservice struct {
					CreatedAt string `json:"created_at"`
					Domains   []struct {
						InternalDomain   string `json:"internal_domain"`
						InternalProtocol string `json:"internal_protocol"`
						PublicPathPrefix string `json:"public_path_prefix"`
					} `json:"domains"`
					Id                      int    `json:"id"`
					IsActive                bool   `json:"is_active"`
					IsVisibleOnAdminSidebar bool   `json:"is_visible_on_admin_sidebar"`
					LogoUrl                 string `json:"logo_url"`
					Name                    string `json:"name"`
					Slug                    string `json:"slug"`
					UpdatedAt               string `json:"updated_at"`
				} `json:"microservice"`
				MicroserviceId int    `json:"microservice_id"`
				Permission     string `json:"permission"`
				UpdatedAt      string `json:"updated_at"`
			} `json:"permissions"`
			Type struct {
				Color     string `json:"color"`
				CreatedAt string `json:"created_at"`
				Id        int    `json:"id"`
				Type      string `json:"type"`
				UpdatedAt string `json:"updated_at"`
			} `json:"type"`
			TypeId    int    `json:"type_id"`
			UpdatedAt string `json:"updated_at"`
		} `json:"roles"`
		SmallAvatarUrl string `json:"small_avatar_url"`
		Theme          string `json:"theme"`
		UpdatedAt      string `json:"updated_at"`
	} `json:"b2b_user"`
	Status int `json:"status"`
}
