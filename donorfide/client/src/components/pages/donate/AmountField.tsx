import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDollarSign } from "@fortawesome/free-solid-svg-icons";
import "./AmountField.sass";
import * as React from "react";

interface amountFieldProps {
	onChange: (newAmount: number) => void
	value: number
}

export default function AmountField(props: amountFieldProps): JSX.Element {
	return <div className="amount-field-wrapper">
		<div className="control has-icons-left has-icons-right amount-field">
			<input className="input is-large" type="number" placeholder="Amount" min={0.50} max={999999.99} step={0.01}
				onChange={(event) => props.onChange(event.target.valueAsNumber * 100)} value={props.value / 100} />
			<span className="icon is-medium is-left">
				<FontAwesomeIcon icon={faDollarSign} />
			</span>
		</div>
	</div>;
}