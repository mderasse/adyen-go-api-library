/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v52/payments ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

// CreatePaymentLinkRequest struct for CreatePaymentLinkRequest
type CreatePaymentLinkRequest struct {
	// List of payments methods to be presented to the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"allowedPaymentMethods\":[\"ideal\",\"giropay\"]`
	AllowedPaymentMethods []string `json:"allowedPaymentMethods,omitempty"`
	Amount                Amount   `json:"amount"`
	BillingAddress        *Address `json:"billingAddress,omitempty"`
	// List of payments methods to be hidden from the shopper. To refer to payment methods, use their `paymentMethod.type` from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: `\"blockedPaymentMethods\":[\"ideal\",\"giropay\"]`
	BlockedPaymentMethods []string `json:"blockedPaymentMethods,omitempty"`
	// The shopper's country code.
	CountryCode     string   `json:"countryCode"`
	DeliveryAddress *Address `json:"deliveryAddress,omitempty"`
	// A short description visible on the Pay By Link page. Maximum length: 280 characters.
	Description string `json:"description,omitempty"`
	// The date that the Pay By Link expires, in ISO 8601 format. For example, 2019-11-23T12:25:28Z. Maximum expiry date should be 30 days from when the payment link is created. If not provided, the default expiry duration is 24 hours.
	ExpiresAt string `json:"expiresAt,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference"`
	// Website URL used for redirection after payment is completed. If provided, a **Continue** button will be shown on the page. If shoppers select the button, they are redirected to the specified URL.
	ReturnUrl string `json:"returnUrl,omitempty"`
	// Indicates whether the payment link can be reused for multiple payments. If not provided, this defaults to **false** which means the link can be used for one successful payment only.
	Reusable bool `json:"reusable,omitempty"`
	// The shopper's email address. We recommend that you provide this data, as it is used in velocity fraud checks. > For 3D Secure 2 transactions, schemes require the `shopperEmail` for both `deviceChannel` **browser** and **app**.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	// The shopper's reference to uniquely identify this shopper (e.g. user ID or account ID). > This field is required for recurring payments.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The physical store, for which this payment is processed.
	Store string `json:"store,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be stored.
	StorePaymentMethod bool `json:"storePaymentMethod,omitempty"`
}
