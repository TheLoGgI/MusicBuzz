import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react';

export const Route = createFileRoute('/musicBuzz')({
    component: AboutComponent,
})

function AboutComponent() {


    <div id="embed-iframe"></div>
    useEffect(() => {
        if ('onSpotifyIframeApiReady' in window)
            window.onSpotifyIframeApiReady = (IFrameAPI: any) => {
                const element = document.getElementById('embed-iframe');
                const options = {
                    uri: 'spotify:episode:7makk4oTQel546B0PZlDM5'
                };
                const callback = (EmbedController: unknown) => {
                    console.log('EmbedController: ', EmbedController);
                };
                IFrameAPI.createController(element, options, callback);
            };
    }, []);


    return (
        <div
            className="hero min-h-screen"
            style={{
                backgroundImage: "url(https://img.daisyui.com/images/stock/photo-1507358522600-9f71e620c44e.jpg)",
            }}>
            <div id="embed-iframe"></div>


        </div>


    )
}
