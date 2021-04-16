import * as React from "react";
import { PaymentRequestButtonElement, useStripe } from "@stripe/react-stripe-js";
import { PaymentRequest, PaymentRequestOptions, StripeError } from "@stripe/stripe-js";
import OrBar from "./OrBar";

interface payNowButtonProps {
	paymentRequest: PaymentRequestOptions
	clientSecret: string

	onError(error: StripeError): void

	onSuccess(): void
}

export default function PayNowButton(props: payNowButtonProps): JSX.Element | null {
	const stripe = useStripe();
	const [paymentRequest, setPaymentRequest] = React.useState(null as PaymentRequest | null);

	const onSuccess = props.onSuccess;
	const onError = props.onError;

	React.useEffect(() => {
		if (stripe) {
			const pr = stripe.paymentRequest(props.paymentRequest);

			pr.canMakePayment().then(result => {
				if (result) {
					pr.on("paymentmethod", async (e) => {
						const { paymentIntent, error: confirmError } = await stripe.confirmCardPayment(
							props.clientSecret,
							{ payment_method: e.paymentMethod.id },
							{ handleActions: false }
						);

						if (confirmError || !paymentIntent) {
							// Something went wrong. Ask the browser for another card
							e.complete("fail");
							return;
						}

						e.complete("success");

						// Check if we need to do something
						if (paymentIntent.status === "requires_action") {
							// Let Stripe.js handle the rest of the payment flow.
							const { error } = await stripe.confirmCardPayment(props.clientSecret);
							if (error) {
								onError(error);
								return;
							}
						}

						onSuccess();
					}
					);
					setPaymentRequest(pr);
				}
			});
		}
	}, [stripe, onError, onSuccess, props.clientSecret, props.paymentRequest]);

	if (paymentRequest) {
		return <>
			<PaymentRequestButtonElement options={{ paymentRequest }} />
			<OrBar />
		</>;
	}

	return null;
}