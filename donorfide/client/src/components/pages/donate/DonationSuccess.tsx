import * as React from "react";
import Hero from "../../ui/Hero";
import { useWindowSize } from "react-use";
import Confetti from "react-confetti";
import Page from "../Page";
import { OrgContext } from "../../App";

export default function DonationSuccess(): JSX.Element {
	const { width, height } = useWindowSize();

	return <OrgContext.Consumer>
		{org => <Page>
			{!org.disable_confetti && <Confetti
				width={width}
				height={height}
			/>}
			<Hero title={org.donation_success_headline || "Thank you!"}
				subtitle={org.donation_success_subheadline || "Your donation has been processed"} center />
			<section className="section">
				<div className="container">
					<p>{org.donation_success_text || "Thank you for your donation. You will receive a receipt via email shortly."}</p>
				</div>
			</section>
		</Page>}
	</OrgContext.Consumer>;

}