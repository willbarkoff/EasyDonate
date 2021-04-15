import {loadStripe, Stripe} from '@stripe/stripe-js'

export let stripe = null as Stripe | null;

export async function prepareStripe(pk: string) {
	stripe = await loadStripe(pk)
}