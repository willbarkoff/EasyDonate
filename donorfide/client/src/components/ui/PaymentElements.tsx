import { CardElement, IbanElement, IdealBankElement } from "@stripe/react-stripe-js";
import * as React from "react";

const style = {
	base: {
		color: "#363636",
		fontFamily: "BlinkMacSystemFont,-apple-system,\"Segoe UI\",Roboto,Oxygen,Ubuntu,Cantarell,\"Fira Sans\",\"Droid Sans\",\"Helvetica Neue\",Helvetica,Arial,sans-serif",
		fontSmoothing: "antialiased",
		fontSize: "16px",
		"::placeholder": {
			color: "#c2c2c2",
		}
	},
	invalid: {
		fontFamily: "BlinkMacSystemFont,-apple-system,\"Segoe UI\",Roboto,Oxygen,Ubuntu,Cantarell,\"Fira Sans\",\"Droid Sans\",\"Helvetica Neue\",Helvetica,Arial,sans-serif",
	}
};

interface paymentElementProps {
	paymentMethod: string
	disabled: boolean
	setPaymentMethod(newMethod: string): void
}


export default function PaymentElement(props: paymentElementProps): JSX.Element {
	function renderPaymentElement() {
		switch (props.paymentMethod) {
			case "card":
				return <CardElement options={{ style, disabled: props.disabled }} />;
			case "iban":
				return <IbanElement options={{ style, supportedCountries: ["SEPA"], disabled: props.disabled }} />;
			case "ideal":
				return <IdealBankElement options={{ style, disabled: props.disabled }} />;
			default:
				return <p>Please choose a payment method above</p>;
		}
	}

	return <div className="field">
		{/*<div className="tabs is-centered">*/}
		{/*	<ul>*/}
		{/*		<li className={props.paymentMethod == "card" ? "is-active" : ""}>*/}
		{/*			<a onClick={() => props.setPaymentMethod("card")}>Card</a>*/}
		{/*		</li>*/}
		{/*		<li className={props.paymentMethod == "iban" ? "is-active" : ""}>*/}
		{/*			<a onClick={() => props.setPaymentMethod("iban")}>SEPA Direct Debit</a>*/}
		{/*		</li>*/}
		{/*		<li className={props.paymentMethod == "ideal" ? "is-active" : ""}>*/}
		{/*			<a onClick={() => props.setPaymentMethod("ideal")}>iDEAL Bank</a>*/}
		{/*		</li>*/}
		{/*		/!*{props.paymentRequest.canMakePayment() && <li className={paymentMethod == "tap" ? "is-active" : ""}>*!/*/}
		{/*		/!*    <a onClick={() => setPaymentMethod("tap")}>Tap to pay</a>*!/*/}
		{/*		/!*</li>}*!/*/}
		{/*	</ul>*/}
		{/*</div>*/}
		<div className="field">
			<label className="label">Payment information</label>
			<div className="control">
				{renderPaymentElement()}
			</div>
		</div>
	</div>;
}
