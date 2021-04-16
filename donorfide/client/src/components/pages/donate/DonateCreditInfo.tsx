import * as React from "react";
import Page from "../Page";
import Hero from "../../ui/Hero";
import {OrgContext} from "../../App";
import {Link, Redirect, useParams} from "react-router-dom";

import "./DonateCreditInfo.sass"
import PaymentElements from "../../ui/PaymentElements";
import {stripe} from "../../../stripe";
import * as Stripe from "@stripe/react-stripe-js"
import {CardElement} from "@stripe/react-stripe-js"
import * as api from "../../../api"
import PayNowButton from "../../ui/PayNowButton";
import {faCircleNotch} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";

interface routeParams {
	amount: string
}

export default function DonateCreditInfo() {
	const {amount} = useParams<routeParams>()
	let amountInt = parseInt(amount, 10)

	const [paymentMethod, setPaymentMethod] = React.useState("card")
	const [isLoading, setIsLoading] = React.useState(false)
	const [isLoadingInitial, setIsLoadingInitial] = React.useState(true)
	const [clientSecret, setClientSecret] = React.useState("")
	const [error, setError] = React.useState(null as string | null | undefined)
	const [email, setEmail] = React.useState("")
	const [name, setName] = React.useState("")
	const [success, setSuccess] = React.useState(false)
	const elements = Stripe.useElements()

	const onSuccess = () => setSuccess(true)

	React.useEffect(() => {
		(async () => {
			let response = await api.POST<api.paymentIntentResponse>("donate/generatePaymentToken", {amount})
			setClientSecret(response.client_secret)
			setIsLoadingInitial(false)
		})();
	}, []);

	async function processPayment() {
		setIsLoading(true)

		if (email == "") {
			setIsLoading(false)
			setError("You must enter an email address.")
			return
		}

		let payload;

		switch (paymentMethod) {
			case "card":
				payload = await stripe!.confirmCardPayment(clientSecret, {
					payment_method: {
						card: elements!.getElement(CardElement)!
					},
					receipt_email: email
				});
				break
			// case "iban":
			// 	payload = await stripe!.confirmSepaDebitPayment(clientSecret, {
			// 		payment_method: {
			// 			sepa_debit: elements!.getElement(IbanElement)!,
			// 			billing_details: {
			// 				name: name,
			// 				email: email
			// 			}
			// 		}
			// 	})
			// 	break
			default:
				setError("That payment method currently isn't supported.");
				setIsLoading(false);
				return
		}

		if (payload.error) {
			setError(payload.error.message);
			setIsLoading(false);
			return
		}

		onSuccess()
	}

	return <OrgContext.Consumer>
		{(org) => {
			const prOpts = {
				country: 'US',
				currency: 'usd',
				total: {label: org.name, amount: amountInt},
				requestPayerName: true,
				requestPayerEmail: true,
			};

			if (success) {
				if (org.donation_success_redirect) {
					// so long!
					document.location.href = org.donation_success_redirect
				}
				return <Redirect to="/donation/success"/>
			}

			return <Page>
				<Hero style="primary" title={org.donate_page_title || "Donate"}
					  subtitle={org.donate_page_subtitle || `Contribute to ${org.name}`} center/>
				<section className="section">
					<div className="container has-text-centered">
						{error && <div className="notification is-danger is-light">{error}</div>}
						<p className="block">
							You're
							donating <strong>${(amountInt / 100).toFixed(2)}</strong> to <strong>{org.name}</strong>.
						</p>
						<p className="block">{org.credit_page_text}</p>
						{
							isLoadingInitial ?
								<p className="block">
									<FontAwesomeIcon icon={faCircleNotch} spin size="3x"/>
								</p>
								: <div className="checkout">
									<div className="field">
										<PayNowButton paymentRequest={prOpts} clientSecret={clientSecret}
													  onError={(e) => setError(e.message)} onSuccess={onSuccess}/>
										<label className="label">Name</label>
										<div className="control">
											<input type="email" value={name} onChange={(e) => setName(e.target.value)}
												   className="input" placeholder="Jane Appleseed" disabled={isLoading}/>
										</div>
									</div>
									<div className="field">
										<label className="label">Email</label>
										<div className="control">
											<input type="email" value={email} onChange={(e) => setEmail(e.target.value)}
												   className="input" placeholder="jane@appleseeds.com"
												   disabled={isLoading}/>
										</div>
									</div>
									<PaymentElements paymentMethod={paymentMethod} setPaymentMethod={setPaymentMethod}
													 disabled={isLoading}/>
									<hr/>
									<div className="buttons is-centered">
										{isLoading ?
											// we need to make a fake button because for some stupid reason you can't disable a link
											<button className="button disabled">Back</button> :
											<Link to="/" className="button">Back</Link>
										}
										<button onClick={processPayment}
												className={`button is-primary ${isLoading ? "is-loading" : ""}`}
												disabled={isLoading}>
											Donate ${(amountInt / 100).toFixed(2)}
										</button>
									</div>
								</div>
						}
					</div>
				</section>
			</Page>
		}}
	</OrgContext.Consumer>
}