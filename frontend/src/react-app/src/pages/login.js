import Head from "next/head"
import Layout from "@/components/base/layout"
import Login from "@/components/page/login"

export default function Login_page(){
    return (
        <>
            <Head>
                <title>シス開</title>
            </Head>
            <Layout>
                <Login />
            </Layout>
        </>
    )
}
