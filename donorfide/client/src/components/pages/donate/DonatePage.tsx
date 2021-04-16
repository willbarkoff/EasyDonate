import * as React from "react";
import Page from "../Page";
import Hero from "../../ui/Hero";
import { OrgContext } from "../../App";
import AmountField from "./AmountField";
import { Link, Redirect } from "react-router-dom";

export default function DonatePage(): JSX.Element {
	const [amount, setAmount] = React.useState(500);

	const [redirect, setRedirect] = React.useState(false);

	function onSubmit(e: React.FormEvent) {
		e.preventDefault();

		setRedirect(true);
	}

	function updateAmount(newAmount: number) {
		if (isNaN(newAmount)) {
			setAmount(0);
			return;
		}

		setAmount(Math.round(newAmount));
	}

	return <OrgContext.Consumer>
		{org => <Page>
			{redirect && <Redirect to={`/donate/${amount}`} />}
			<Hero style="primary" title={org.donate_page_title || "Donate"}
				subtitle={org.donate_page_subtitle || `Contribute to ${org.name}`} center />
			<section className="section">
				<form onSubmit={onSubmit}>
					{org.home_page_text && <p className="block has-text-centered">{org.home_page_text}</p>}
					<div className="container has-text-centered">
						<AmountField onChange={updateAmount} value={amount} />
						{/* Use an input so that the browser can handle validation. */}
						<input type="submit" className="button is-primary m-5" value="Continue" />
					</div>
				</form>
			</section>
		</Page>}
	</OrgContext.Consumer>;
}