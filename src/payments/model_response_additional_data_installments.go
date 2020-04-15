/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v52/authorise ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments

// ResponseAdditionalDataInstallments struct for ResponseAdditionalDataInstallments
type ResponseAdditionalDataInstallments struct {
	// The number of installments that the payment amount should be charged with.  Example: 5 > Only relevant for card payments in countries that support installments.
	InstallmentsValue string `json:"installments.value,omitempty"`
	// Type of installment. The value of `installmentType` should be **IssuerFinanced**.
	InstallmentPaymentDataInstallmentType string `json:"installmentPaymentData.installmentType,omitempty"`
	// Possible values: * PayInInstallmentsOnly * PayInFullOnly * PayInFullOrInstallments
	InstallmentPaymentDataPaymentOptions string `json:"installmentPaymentData.paymentOptions,omitempty"`
	// Total number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrNumberOfInstallments string `json:"installmentPaymentData.option[itemNr].numberOfInstallments,omitempty"`
	// Interest rate for the installment period.
	InstallmentPaymentDataOptionItemNrInterestRate string `json:"installmentPaymentData.option[itemNr].interestRate,omitempty"`
	// Installment fee amount in minor units.
	InstallmentPaymentDataOptionItemNrInstallmentFee string `json:"installmentPaymentData.option[itemNr].installmentFee,omitempty"`
	// Annual interest rate.
	InstallmentPaymentDataOptionItemNrAnnualPercentageRate string `json:"installmentPaymentData.option[itemNr].annualPercentageRate,omitempty"`
	// First Installment Amount in minor units.
	InstallmentPaymentDataOptionItemNrFirstInstallmentAmount string `json:"installmentPaymentData.option[itemNr].firstInstallmentAmount,omitempty"`
	// Subsequent Installment Amount in minor units.
	InstallmentPaymentDataOptionItemNrSubsequentInstallmentAmount string `json:"installmentPaymentData.option[itemNr].subsequentInstallmentAmount,omitempty"`
	// Minimum number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrMinimumNumberOfInstallments string `json:"installmentPaymentData.option[itemNr].minimumNumberOfInstallments,omitempty"`
	// Maximum number of installments possible for this payment.
	InstallmentPaymentDataOptionItemNrMaximumNumberOfInstallments string `json:"installmentPaymentData.option[itemNr].maximumNumberOfInstallments,omitempty"`
	// Total amount in minor units.
	InstallmentPaymentDataOptionItemNrTotalAmountDue string `json:"installmentPaymentData.option[itemNr].totalAmountDue,omitempty"`
}
