import * as React from "react";
import Page from "../Page";
import Hero from "../../ui/Hero";
import {OrgContext} from "../../App";
import AmountField from "./AmountField";
import {Link} from "react-router-dom";

export default function DonatePage() {
	const [amount, setAmount] = React.useState(500)

	return <OrgContext.Consumer>
		{org => <Page>
			<Hero style="primary" title={org.donate_page_title || "Donate"}
				  subtitle={org.donate_page_subtitle || `Contribute to ${org.name}`} center/>
			<section className="section">
				{org.home_page_text && <p className="block has-text-centered">{org.home_page_text}</p>}
				<div className="container has-text-centered">
					<AmountField onChange={setAmount} value={amount}/>
					<Link to={`/donate/${amount}`} className="button is-primary m-5">Continue</Link>
				</div>
			</section>
		</Page>}
	</OrgContext.Consumer>
}