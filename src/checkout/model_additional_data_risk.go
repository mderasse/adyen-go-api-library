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
// AdditionalDataRisk struct for AdditionalDataRisk
type AdditionalDataRisk struct {
	// The data for your custom risk field. For more information, refer to [Create custom risk fields](https://docs.adyen.com/risk-management/configure-custom-risk-rules#step-1-create-custom-risk-fields).
	RiskdataCustomFieldName string `json:"riskdata.[customFieldName],omitempty"`
	// The price of item in the basket, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataBasketItemItemNrAmountPerItem string `json:"riskdata.basket.item[itemNr].amountPerItem,omitempty"`
	// Brand of the item.
	RiskdataBasketItemItemNrBrand string `json:"riskdata.basket.item[itemNr].brand,omitempty"`
	// Category of the item.
	RiskdataBasketItemItemNrCategory string `json:"riskdata.basket.item[itemNr].category,omitempty"`
	// Color of the item.
	RiskdataBasketItemItemNrColor string `json:"riskdata.basket.item[itemNr].color,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataBasketItemItemNrCurrency string `json:"riskdata.basket.item[itemNr].currency,omitempty"`
	// ID of the item.
	RiskdataBasketItemItemNrItemID string `json:"riskdata.basket.item[itemNr].itemID,omitempty"`
	// Manufacturer of the item.
	RiskdataBasketItemItemNrManufacturer string `json:"riskdata.basket.item[itemNr].manufacturer,omitempty"`
	// A text description of the product the invoice line refers to.
	RiskdataBasketItemItemNrProductTitle string `json:"riskdata.basket.item[itemNr].productTitle,omitempty"`
	// Quantity of the item purchased.
	RiskdataBasketItemItemNrQuantity string `json:"riskdata.basket.item[itemNr].quantity,omitempty"`
	// Email associated with the given product in the basket (usually in electronic gift cards).
	RiskdataBasketItemItemNrReceiverEmail string `json:"riskdata.basket.item[itemNr].receiverEmail,omitempty"`
	// Size of the item.
	RiskdataBasketItemItemNrSize string `json:"riskdata.basket.item[itemNr].size,omitempty"`
	// [Stock keeping unit](https://en.wikipedia.org/wiki/Stock_keeping_unit).
	RiskdataBasketItemItemNrSku string `json:"riskdata.basket.item[itemNr].sku,omitempty"`
	// [Universal Product Code](https://en.wikipedia.org/wiki/Universal_Product_Code).
	RiskdataBasketItemItemNrUpc string `json:"riskdata.basket.item[itemNr].upc,omitempty"`
	// Code of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionCode string `json:"riskdata.promotions.promotion[itemNr].promotionCode,omitempty"`
	// The discount amount of the promotion, represented in [minor units](https://docs.adyen.com/development-resources/currency-codes).
	RiskdataPromotionsPromotionItemNrPromotionDiscountAmount string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountAmount,omitempty"`
	// The three-character [ISO currency code](https://en.wikipedia.org/wiki/ISO_4217).
	RiskdataPromotionsPromotionItemNrPromotionDiscountCurrency string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountCurrency,omitempty"`
	// Promotion's percentage discount. It is represented in percentage value and there is no need to include the '%' sign.  e.g. for a promotion discount of 30%, the value of the field should be 30.
	RiskdataPromotionsPromotionItemNrPromotionDiscountPercentage string `json:"riskdata.promotions.promotion[itemNr].promotionDiscountPercentage,omitempty"`
	// Name of the promotion.
	RiskdataPromotionsPromotionItemNrPromotionName string `json:"riskdata.promotions.promotion[itemNr].promotionName,omitempty"`
	// Reference number of the risk profile that you want to apply to the payment. If not provided or left blank, the merchant-level account's default risk profile will be applied to the payment. For more information, see [dynamically assign a risk profile to a payment](https://docs.adyen.com/risk-management/create-and-use-risk-profiles#dynamically-assign-a-risk-profile-to-a-payment).
	RiskdataRiskProfileReference string `json:"riskdata.riskProfileReference,omitempty"`
}
