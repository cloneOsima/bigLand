import React, { useEffect } from "react";

function Map() {
    useEffect(() => {
        // 이미 스크립트가 있으면 중복 추가 방지
        const scriptId = "kakao-map-script";
        const existingScript = document.getElementById(scriptId);

        function onLoadKakaoMap() {
            const container = document.getElementById('map');
            const options = {
                center: new window.kakao.maps.LatLng(37.5665, 126.9780),
                level: 3,
            };
            new window.kakao.maps.Map(container, options);
        }

        if (!existingScript) {
            const script = document.createElement("script");
            script.id = scriptId;
            script.src = "//dapi.kakao.com/v2/maps/sdk.js?appkey=b642534716be93d3829d4d9c9fa98dcc";
            script.async = true;
            script.onload = onLoadKakaoMap;
            document.head.appendChild(script);
        } else if (window.kakao && window.kakao.maps) {
            onLoadKakaoMap();
        } else {
            existingScript.onload = onLoadKakaoMap;
        }
    }, []);

    return (
        <div id="map" style={{ width: "500px", height: "400px" }}></div>
    );
}

export default Map;