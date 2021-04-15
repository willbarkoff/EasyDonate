import * as React from "react";
import Page from "../Page";
import Hero from "../../ui/Hero";
import {OrgContext} from "../../App";
import {Link, useParams} from "react-router-dom";
import {stripe} from "../../../stripe";
import {CardElement, Elements} from "@stripe/react-stripe-js";

import "./DonateCreditInfo.sass"

let style = {
	base: {
		color: "#363636",
		fontFamily: 'BlinkMacSystemFont,-apple-system,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",Helvetica,Arial,sans-serif',
		fontSmoothing: "antialiased",
		fontSize: "16px",
		"::placeholder": {
			color: "#c2c2c2",
		}
	},
	invalid: {
		fontFamily: 'BlinkMacSystemFont,-apple-system,"Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",Helvetica,Arial,sans-serif',
	}
}

interface routeParams {
	amount: string
}

export default function DonateCreditInfo() {
	const {amount} = useParams<routeParams>()
	let amountInt = parseInt(amount, 10)

	return <OrgContext.Consumer>
		{org => <Page>
			<Hero style="primary" title={org.donate_page_title || "Donate"}
				  subtitle={org.donate_page_subtitle || `Contribute to ${org.name}`} center/>
			<section className="section">
				<div className="container has-text-centered">
					<p>
						You're
						donating <strong>${(amountInt / 100).toFixed(2)}</strong> to <strong>{org.name}</strong>.
					</p>
					<div className="checkout">
						<div className="field">
							<label className="label">Email</label>
							<div className="control">
								<input type="email" className="input" placeholder="Your email address"/>
							</div>
						</div>
						<hr/>
						<Elements stripe={stripe}>
							<div className="field">
								<label className="label">Credit card</label>
								<CardElement options={{style: style}}/>
							</div>
							<Link to="/" className="button">Back</Link>
						</Elements>
					</div>
				</div>
			</section>
		</Page>}
	</OrgContext.Consumer>
}