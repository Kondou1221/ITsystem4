import Header from "@/components/base/header";

export default function Layout({ children }) {
    return (
        <>
            <Header />
            <main>{children}</main>
        </>
    );
}