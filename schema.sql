-- 1. USERS
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       email TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       full_name TEXT NOT NULL,
                       phone_number TEXT NOT NULL ,
                       birth_date DATE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 2. ADDRESSES
CREATE TABLE addresses (
                           id BIGSERIAL PRIMARY KEY,
                           user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                           label TEXT, -- e.g. "Home", "Office"
                           line1 TEXT NOT NULL,
                           line2 TEXT,
                           city TEXT NOT NULL,
                           postal_code TEXT NOT NULL,
                           country_code CHAR(2) NOT NULL,
                           is_default BOOLEAN DEFAULT FALSE,
                           created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                           updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 3. CATEGORIES
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name TEXT NOT NULL,
                            slug TEXT UNIQUE NOT NULL,
                            parent_id INT REFERENCES categories(id)
);

-- 4. PRODUCTS
CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          category_id INT REFERENCES categories(id),
                          name TEXT NOT NULL,
                          description TEXT,
                          base_price DECIMAL(10, 2) NOT NULL,
                          image_url TEXT NOT NULL,
                          is_active BOOLEAN DEFAULT TRUE,
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 5. PRODUCT VARIANTS (Inventory & Specifics)
CREATE TABLE product_variants (
                                  id BIGSERIAL PRIMARY KEY,
                                  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
                                  sku TEXT UNIQUE NOT NULL,
                                  size TEXT NOT NULL,
                                  color TEXT NOT NULL,
                                  stock_quantity INT DEFAULT 0 CHECK (stock_quantity >= 0),
                                  price_override DECIMAL(10, 2), -- Overrides base_price if set
                                  discounted_price DECIMAL(10, 2), -- If set, this is the "Sale" price
                                  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 6. LOOKS (Outfits)
CREATE TABLE looks (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       description TEXT,
                       image_url TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Junction Table: Look <-> Product
CREATE TABLE look_products (
                               look_id INT REFERENCES looks(id) ON DELETE CASCADE,
                               product_id BIGINT REFERENCES products(id) ON DELETE CASCADE,
                               PRIMARY KEY (look_id, product_id)
);

-- 7. CARTS
CREATE TABLE carts (
                       id BIGSERIAL PRIMARY KEY,
                       user_id BIGINT REFERENCES users(id) ON DELETE SET NULL, -- Nullable for Guest
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE cart_items (
                            id BIGSERIAL PRIMARY KEY,
                            cart_id BIGINT REFERENCES carts(id) ON DELETE CASCADE,
                            product_variant_id BIGINT REFERENCES product_variants(id),
                            quantity INT NOT NULL CHECK (quantity > 0),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                            UNIQUE(cart_id, product_variant_id) -- Prevent duplicate rows for same item
);

-- 8. ORDERS
CREATE TABLE orders (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT REFERENCES users(id),
                        total_amount DECIMAL(10, 2) NOT NULL,
                        total_discount_saved DECIMAL(10, 2) DEFAULT 0, -- Stored specifically for history
                        status TEXT NOT NULL DEFAULT 'PENDING', -- PENDING, PAID, SHIPPED, DELIVERED, CANCELLED
                        shipping_address_json JSONB NOT NULL,
                        payment_provider_id TEXT, -- Stripe Intent ID
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE order_items (
                             id BIGSERIAL PRIMARY KEY,
                             order_id BIGINT REFERENCES orders(id) ON DELETE CASCADE,
                             product_variant_id BIGINT REFERENCES product_variants(id),
                             quantity INT NOT NULL,
                             price_at_purchase DECIMAL(10, 2) NOT NULL -- Snapshot of price
);