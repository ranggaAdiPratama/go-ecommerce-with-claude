import { Header } from '../components/layout/Header';
import { Footer } from '../components/layout/Footer';
import { ShopList } from '../components/shops/ShopList';
import { usePageTitle } from '../hooks/usePageTitle';
import { CategoryList } from '../components/categories/CategoryList';

export const HomePage = () => {
    usePageTitle()

    return (
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
            <Header />
            <main>
                <CategoryList />
                <ShopList />
            </main>
            <Footer />
        </div>
    );
};