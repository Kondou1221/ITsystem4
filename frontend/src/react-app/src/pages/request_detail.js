import Head from "next/head"
import Layout from "@/components/base/layout"
import Request_detail from "@/components/page/request_detail"

export default function Request_detail_page(){
    return (
        <>
            <Head>
                <title>シス開</title>
            </Head>
            <Layout>
                <Request_detail />
            </Layout>
        </>
    )
}
