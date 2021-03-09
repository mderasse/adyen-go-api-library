/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 67
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// StoredPaymentMethod struct for StoredPaymentMethod
type StoredPaymentMethod struct {
	// The brand of the card.
	Brand string `json:"brand,omitempty"`
	// The month the card expires.
	ExpiryMonth string `json:"expiryMonth,omitempty"`
	// The year the card expires.
	ExpiryYear string `json:"expiryYear,omitempty"`
	// The unique payment method code.
	HolderName string `json:"holderName,omitempty"`
	// The IBAN of the bank account.
	Iban string `json:"iban,omitempty"`
	// A unique identifier of this stored payment method.
	Id string `json:"id,omitempty"`
	// The last four digits of the PAN.
	LastFour string `json:"lastFour,omitempty"`
	// The display name of the stored payment method.
	Name string `json:"name,omitempty"`
	// The name of the bank account holder.
	OwnerName string `json:"ownerName,omitempty"`
	// The shopper’s email address.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The supported shopper interactions for this stored payment method.
	SupportedShopperInteractions []string `json:"supportedShopperInteractions,omitempty"`
	// The type of payment method.
	Type string `json:"type,omitempty"`
}
