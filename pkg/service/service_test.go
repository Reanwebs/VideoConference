package service

// import (
// 	"conference/pkg/common/config"
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang/mock/gomock"
// )

// func Test_UserSignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)

// 	// Create mock implementations for the repositories
// 	userRepo := mockrepo.NewMockUserRepository(ctrl)
// 	orderRepo := mockrepo.NewMockOrderRepository(ctrl)
// 	otpRepo := mockrepo.NewMockOtpRepository(ctrl)
// 	inventoryRepo := mockrepo.NewMockInventoryRepository(ctrl)
// 	helper := mockhelper.NewMockHelper(ctrl)
// 	cfg := config.Config{}

// 	userUseCase := NewUserUseCase(userRepo, cfg, otpRepo, inventoryRepo, orderRepo, helper)

// 	testData := map[string]struct {
// 		input          models.UserDetails
// 		StubDetails    func(mockrepo.MockUserRepository, mockrepo.MockOrderRepository, models.UserDetails, mockhelper.MockHelper)
// 		expectedOutput models.TokenUsers
// 		expectedError  error
// 	}{
// 		"success": {
// 			input: models.UserDetails{
// 				Name:            "Arun K",
// 				Email:           "arthurbishop120@gmail.com",
// 				Phone:           "6282246077",
// 				Password:        gomock.Any().String(),
// 				ConfirmPassword: gomock.Any().String(),
// 			},
// 			StubDetails: func(userRepo mockrepo.MockUserRepository, orderRepo mockrepo.MockOrderRepository, signupData models.UserDetails, helper mockhelper.MockHelper) {
// 				gomock.InOrder(
// 					userRepo.EXPECT().CheckUserAvailability(signupData.Email).Times(1).Return(false),
// 					userRepo.EXPECT().FindUserFromReference("12345").Times(1).Return(1, nil),
// 					helper.EXPECT().PasswordHashing(signupData.Password).Times(1).Return(gomock.Any().String(), nil),
// 					helper.EXPECT().GenerateRefferalCode().Times(1).Return(gomock.Any().String(), nil),
// 					userRepo.EXPECT().UserSignUp(signupData, gomock.Any().String()).Times(1).Return(
// 						models.UserDetailsResponse{
// 							Id:    1,
// 							Name:  signupData.Name,
// 							Email: signupData.Email,
// 							Phone: signupData.Phone,
// 						}, nil,
// 					),
// 					helper.EXPECT().GenerateTokenClients(models.UserDetailsResponse{
// 						Id:    1,
// 						Name:  "Arun K",
// 						Email: "arthurbishop120@gmail.com",
// 						Phone: "6282246077",
// 					}).Times(1).Return(gomock.Any().String(), nil),
// 					userRepo.EXPECT().CreditReferencePointsToWallet(1).Times(1).Return(nil),
// 					orderRepo.EXPECT().CreateNewWallet(1).Times(1).Return(1, nil),
// 				)
// 			},
// 			expectedOutput: models.TokenUsers{
// 				Users: models.UserDetailsResponse{
// 					Id:    1,
// 					Name:  "Arun K",
// 					Email: "arthurbishop120@gmail.com",
// 					Phone: "6282246077",
// 				},
// 				Token: "ajjsjsjsjsjs.sjsjsjsjsjs.sjsjsjsjs",
// 			},
// 			expectedError: nil,
// 		},
// 	}

// 	for _, test := range testData {

// 		test.StubDetails(*userRepo, *orderRepo, test.input, *helper)

// 		tokenusers, err := userUseCase.UserSignUp(test.input, "12345")

// 		assert.Equal(t, test.expectedOutput.Users.Id, tokenusers.Users.Id)
// 		assert.Equal(t, test.expectedOutput.Users.Name, tokenusers.Users.Name)
// 		assert.Equal(t, test.expectedOutput.Users.Email, tokenusers.Users.Email)
// 		assert.Equal(t, test.expectedError, err)

// 	}
// }
