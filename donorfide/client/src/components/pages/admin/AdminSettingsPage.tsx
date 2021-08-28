import * as React from "react";
import Page from "../Page";
import Hero from "../../ui/Hero";
import { Link } from "react-router-dom";
import AdminStringSetting from "./AdminStringSetting";
import { OrgContext } from "../../App";

export default function AdminSettingsPage(): JSX.Element {
	return <OrgContext.Consumer>
		{org => <Page>
			<Hero title="Donorfide Settings" />
			<section className="section">
				<div className="container">
					<div className="buttons">
						<Link to="/admin" className="button">Back</Link>
					</div>
					<p className="block">This page allows you to change how Donorfide looks and behaves for your users.
						Note
						that you may need to refresh the page to see your changes.</p>
					<AdminStringSetting originalValue={org.name} settingKey={"orgName"} label={"Name"}
						description={"The name of your organization"} />
					<AdminStringSetting originalValue={org.stripe_pk} settingKey={"stripePK"}
						label={"Stripe publishable key"}
						description={"The publishable key of your stripe account."} />
					<AdminStringSetting originalValue={""} settingKey={"stripeSK"} label={"Stripe secret key"}
						description={"The secret key of your Stripe account."} />
					<AdminStringSetting originalValue={org.home_page_text} settingKey={"homePageText"}
						label={"Home page text"}
						description={"The text above the amount field on the home page."} />
					<AdminStringSetting originalValue={org.receipt_text} settingKey={"chargeDescription"} label={"Receipt text"} description={"Text to include on the customer receipt. This is usually infomration about the organization's tax-deduction status."} />
					<AdminStringSetting originalValue={org.donate_page_title} settingKey={"donatePageTitle"}
						label={"Donate page title"}
						description={"The title of the donation page. By default this is the word \"Donate\""} />
					<AdminStringSetting originalValue={org.donate_page_subtitle} settingKey={"donatePageSubtitle"}
						label={"Donate page subtitle"}
						description={`The subtitle of the donation page. By default this is the phrase "Contribute to ${org.name}."`} />
					<AdminStringSetting originalValue={org.credit_page_text} settingKey={"creditPageText"}
						label={"Credit card page text"}
						description={"This text goes on the page where users enter their credit card information."} />
					<AdminStringSetting originalValue={org.imprint} settingKey={"orgImprint"} label={"Imprint"}
						description={"The imprint on the footer of each page"} />
					<AdminStringSetting originalValue={org.donation_success_headline}
						settingKey={"donationSuccessHeadline"} label={"Donation success headline"}
						description={"The headline for the donation success page. By default, this is the phrase \"Thank you!\""} />
					<AdminStringSetting originalValue={org.donation_success_subheadline}
						settingKey={"donationSuccessSubheadline"} label={"Donation success subheadline"}
						description={"The subheadline for the donation success page. By default, this is the phrase \"Your donation has been processed.\""} />
					<AdminStringSetting originalValue={org.donation_success_text} settingKey={"donationSuccessText"}
						label={"Donation success text"}
						description={"The text displayed on the donation success page. By default, this is the phrase \"Thank you for your donation. You will receive a receipt via email shortly.\""} />
					<AdminStringSetting originalValue={org.donation_success_redirect}
						settingKey={"donationSuccessRedirect"} label={"Donation success redirect"}
						description={"The page to be redirected to upon donation completion. Make sure to include the \"http\" or \"https\". To use the default Donorfide page, leave this setting blank."} />
				</div>
			</section>
		</Page>}
	</OrgContext.Consumer>;
}