package tokenx

import (
	"testing" // Importing the testing package for writing test cases
	"time"    // Importing the time package to handle durations

	"github.com/Telktia-LTD/longswipe-reuse/interfacesx" // Importing interfaces from the project's x/interfacesx package

	"github.com/stretchr/testify/require" // Importing testify's require package for assertions in tests
)

func TestPasetoMaker(t *testing.T) {
	// Define a symmetric key
	symmetricKey := "0123456789abcdef0123456789abcdef" // Defining a 32-byte symmetric key

	// Create a new PasetoMaker
	maker, err := NewPasetoMaker(symmetricKey) // Creating a new instance of PasetoMaker with the symmetric key
	require.NoError(t, err)                    // Asserting that there is no error during creation
	require.NotEmpty(t, maker)                 // Asserting that the maker instance is not empty

	// Define a user and token duration
	user := interfacesx.UserResponse{
		Username: "testuser", // Creating a user response with username "testuser"
	}
	duration := time.Minute // Setting the token duration to one minute

	// Create a token
	token, payload, err := maker.CreateToken(user, duration) // Creating a token for the user with the specified duration
	require.NoError(t, err)                                  // Asserting that there is no error during token creation
	require.NotEmpty(t, token)                               // Asserting that the created token is not empty
	require.NotNil(t, payload)                               // Asserting that the payload is not nil

	// Verify the token
	verifiedPayload, err := maker.VerifyToken(token) // Verifying the created token
	require.NoError(t, err)                          // Asserting that there is no error during token verification
	require.NotNil(t, verifiedPayload)               // Asserting that the verified payload is not nil

	// Check if the payload matches
	require.Equal(t, payload.User.Username, verifiedPayload.User.Username)               // Asserting that the username in the payload matches the verified payload
	require.WithinDuration(t, payload.IssuedAt, verifiedPayload.IssuedAt, time.Second)   // Asserting that the issued time is within a second of the verified issued time
	require.WithinDuration(t, payload.ExpiresAt, verifiedPayload.ExpiresAt, time.Second) // Asserting that the expiry time is within a second of the verified expiry time
}

func TestExpiredPasetoToken(t *testing.T) {
	symmetricKey := "0123456789abcdef0123456789abcdef" // Defining a 32-byte symmetric key

	maker, err := NewPasetoMaker(symmetricKey) // Creating a new instance of PasetoMaker with the symmetric key
	require.NoError(t, err)                    // Asserting that there is no error during creation
	require.NotEmpty(t, maker)                 // Asserting that the maker instance is not empty

	user := interfacesx.UserResponse{
		Username: "testuser", // Creating a user response with username "testuser"
	}
	duration := -time.Minute // Setting the token duration to a negative value to simulate an expired token

	token, payload, err := maker.CreateToken(user, duration) // Creating a token for the user with the expired duration
	require.NoError(t, err)                                  // Asserting that there is no error during token creation
	require.NotEmpty(t, token)                               // Asserting that the created token is not empty
	require.NotNil(t, payload)                               // Asserting that the payload is not nil

	verifiedPayload, err := maker.VerifyToken(token)    // Verifying the expired token
	require.Error(t, err)                               // Asserting that there is an error during token verification
	require.Nil(t, verifiedPayload)                     // Asserting that the verified payload is nil
	require.EqualError(t, err, ErrExpiredToken.Error()) // Asserting that the error is an expired token error
}

func TestInvalidPasetoToken(t *testing.T) {
	symmetricKey := "0123456789abcdef0123456789abcdef" // Defining a 32-byte symmetric key

	maker, err := NewPasetoMaker(symmetricKey) // Creating a new instance of PasetoMaker with the symmetric key
	require.NoError(t, err)                    // Asserting that there is no error during creation
	require.NotEmpty(t, maker)                 // Asserting that the maker instance is not empty

	user := interfacesx.UserResponse{
		Username: "testuser", // Creating a user response with username "testuser"
	}
	duration := time.Minute // Setting the token duration to one minute

	token, payload, err := maker.CreateToken(user, duration) // Creating a token for the user with the specified duration
	require.NoError(t, err)                                  // Asserting that there is no error during token creation
	require.NotEmpty(t, token)                               // Asserting that the created token is not empty
	require.NotNil(t, payload)                               // Asserting that the payload is not nil

	// Modify the token to make it invalid
	token = token[:len(token)-1] // Removing the last character of the token to make it invalid

	verifiedPayload, err := maker.VerifyToken(token)    // Verifying the invalid token
	require.Error(t, err)                               // Asserting that there is an error during token verification
	require.Nil(t, verifiedPayload)                     // Asserting that the verified payload is nil
	require.EqualError(t, err, ErrInvalidToken.Error()) // Asserting that the error is an invalid token error
}
