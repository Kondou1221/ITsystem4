import Head from "next/head"
import Layout from "@/components/base/layout"
import Request_list from "@/components/page/request_list"
import Searchrequest from "@/features/request/api/getallRequest"

export default function Home(props){
    return (
        <>
            <Head>
                <title>シス開</title>
            </Head>
            <Layout>
                <Request_list {...props.request_data} />
            </Layout>
        </>
    )
}

export async function getServerSideProps() {
    // try {
    //     const products = await fetch("http://back_go:8080/request_get")
    //         .then(data => data.json())
    //     return {
    //         props: {
    //             products,
    //         }
    //     }
    // } catch (e) {
    //     console.log(e)
    //     return {
    //         props: {
    //             products: [e.message],
    //         }
    //     }
    // }
    const res = await Searchrequest()
    return {
        props: {
            request_data: res
        }
    }
}
