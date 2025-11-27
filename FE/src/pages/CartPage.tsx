import { usePageTitle } from "../hooks/usePageTitle"

export const CartPage = () => {
    usePageTitle("Cart")

    return (
        <div>
            <h1>Shopping Cart</h1>
        </div>
    )
}