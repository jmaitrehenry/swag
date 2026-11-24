package main

import "net/http"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for exercising formData request bodies.
func main() {
	http.HandleFunc("/import", ImportSettings)
	http.HandleFunc("/checks/update", UpdateCheck)
}

// @Summary Import settings
// @Accept          multipart/form-data
// @Produce         plain
// @Param           zip                   formData file    true  "Zip file containing checks.toml, users.toml, and/or settings.toml"
// @Param           destructiveChecks     formData boolean false "If true, destructively replace existing checks during checks.toml import"
// @Param           destructiveUsers      formData boolean false "If true, destructively replace existing users during users.toml import"
// @Param           destructiveSettings   formData boolean false "If true, destructively replace existing settings during settings.toml import"
// @Success 200
// @Router /import [post]
func ImportSettings(w http.ResponseWriter, r *http.Request) {
}

// @Summary Update check
// @Accept          application/x-www-form-urlencoded
// @Param           id          formData    int     true    "Check ID to update"
// @Param           name        formData    string  false   "Update check name"
// @Param           description formData    string  false   "Update check description"
// @Param           weight      formData    int     false   "Update check weight"
// @Param           activated   formData    boolean false   "Update check activation"
// @Param           allow_user_secrets  formData boolean false  "Are teams allowed to change the checks auth parameters?"
// @Param           type        formData    string  false   "Update check type"
// @Param           source      formData    string  false   "Update check source"
// @Param           metadata    formData    string  false   "Update check metadata"
// @Success 200
// @Router /checks/update [post]
func UpdateCheck(w http.ResponseWriter, r *http.Request) {
}
