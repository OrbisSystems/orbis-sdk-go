package model

type LoginByEmailRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	DeviceID   string `json:"device_id"`
	RememberMe bool   `json:"remember_me"`
}

type LoginByEmailResponse struct {
	Status     int `json:"status"`
	LoginBasic struct {
		Tokens struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			PairId       string `json:"pair_id"`
		} `json:"tokens"`
	} `json:"login_basic"`
}

type LoginByAPIKeyRequest struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type LoginByAPIKeyResponse struct {
	Status       int `json:"status"`
	ApiKeysLogin struct {
		Tokens struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			PairId       string `json:"pair_id"`
		} `json:"tokens"`
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
