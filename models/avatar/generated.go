package avatar

// Authorization
//
// POST {{api}}/authorize?env={{env}}
//
// Get the token for subsequent requests
type Authorization struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

// Renew
//
// POST {{api}}/renew
//
// Renew a generated token
type Renew struct {
	Token string `json:"token"`
}

// Invalidate
//
// POST {{api}}/invalidate
//
// Invalidate a generated token
type Invalidate struct {
	Message string `json:"message"`
}

// Upload user avatar
//
// POST {{api}}/user/:userId?token={{token}}
type UploadUserAvatar struct {
	Message string `json:"message"`
}

// Remove user avatar
//
// DELETE {{api}}/user/:userId?token={{token}}
type RemoveUserAvatar struct {
	Message string `json:"message"`
}

// Upload account avatar
//
// POST {{api}}/user/:accountId?token={{token}}
type UploadAccountAvatar struct {
	Message string `json:"message"`
}

// Remove account avatar
//
// DELETE {{api}}/user/:accountId?token={{token}}
type RemoveAccountAvatar struct {
	Message string `json:"message"`
}
