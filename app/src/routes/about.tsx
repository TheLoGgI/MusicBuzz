import { createFileRoute } from '@tanstack/react-router'
import { getData } from '../hooks/data/getData';

export const Route = createFileRoute('/about')({
    component: AboutComponent,
})

function AboutComponent() {
    const { isPending, error, data } = getData()
    console.log('data: ', data);
    return (
        <div
            className="hero min-h-screen"
            style={{
                backgroundImage: "url(https://img.daisyui.com/images/stock/photo-1507358522600-9f71e620c44e.jpg)",
            }}>
            <div className="hero-overlay bg-opacity-60"></div>
            <div className="hero-content text-neutral-content text-center">
                <div className="max-w-md space-y-4">
                    <h1 className="mb-5 text-5xl font-bold">Hello there</h1>
                    {isPending ? (
                        <p>Loading...</p>
                    ) : error ? (
                        <p>Error: {error.message}</p>
                    ) : (
                        <p className='mx-auto w-3/4'>{data.description} </p>
                    )}
                    <button className="btn btn-primary">Get Started</button>
                </div>
            </div>
        </div>


    )
}
