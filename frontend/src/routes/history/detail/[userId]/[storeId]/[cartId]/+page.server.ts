import type { storeOrderInfo } from '$lib';
import { BACKEND_PATH as backendPath } from '$env/static/private';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return {
		cartInfos: await getCart(params.userId, params.cartId, params.storeId)
	};
};
async function getCart(userId: string, cartId: string, storeId: string) {
	try {
		const resp = await fetch(
			backendPath + '/customer/' + userId + '/cart/' + cartId + '/store/' + storeId + '/carts'
		);
		if (resp.ok) {
			return (await resp.json()) as storeOrderInfo;
		}
		return (await (await fetch('/updateOrderDetail.json')).json()) as storeOrderInfo;
	} catch {
		return (await (await fetch('/updateOrderDetail.json')).json()) as storeOrderInfo;
	}
}
