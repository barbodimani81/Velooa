"I am building a clothing e-commerce site. I have a custom Go backend API and I need you to build the Frontend components using React (Next.js) and Tailwind CSS.

Key Business Logic:

Variants: Products are not simple. A product (e.g., 'T-Shirt') has multiple variants (Size/Color). The UI must force the user to select a Size/Color before the 'Add to Cart' button is enabled.

Looks: We have 'Looks' (Outfits). The Look Detail page shows a main outfit photo, and below it, a list of the individual products in that outfit. The user can add specific items from the look to the cart.

Pricing: Show both original_price and discounted_price if a discount exists. Highlight the savings.

Data Structures:

Cart Item: { product_name, variant_info: "Red / M", quantity, unit_price, total_price }

Product: { id, name, images[], base_price, variants: [{id, size, color, stock}] }

Task: Please create the Product Detail Page component. It should have an image gallery on the left, and on the right: Title, Price, a selector for Color, a selector for Size (disable sizes with 0 stock), and an Add to Cart button. Use mock data that matches the structure above."