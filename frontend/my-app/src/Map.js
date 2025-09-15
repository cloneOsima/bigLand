import React, { useEffect, useRef } from 'react';

const Map = () => {
  const mapElement = useRef(null);
  const mapInstance = useRef(null);

  useEffect(() => {
    // 네이버 지도 API가 로드되었는지 확인
    if (!window.naver || !window.naver.maps) {
      console.error('네이버 지도 API가 로드되지 않았습니다.');
      return;
    }

    // 지도 옵션 설정
    const mapOptions = {
      center: new window.naver.maps.LatLng(37.5665, 126.9780), // 서울 시청 좌표
      zoom: 15,
      mapTypeId: window.naver.maps.MapTypeId.NORMAL,
    };

    // 지도 생성
    mapInstance.current = new window.naver.maps.Map(mapElement.current, mapOptions);

    // 마커 추가 (선택사항)
    const marker = new window.naver.maps.Marker({
      position: new window.naver.maps.LatLng(37.5665, 126.9780),
      map: mapInstance.current,
      title: '서울 시청'
    });

    // 정보 창 추가 (선택사항)
    const infoWindow = new window.naver.maps.InfoWindow({
      content: '<div style="padding:10px;min-width:200px;line-height:150%;"><h4>서울특별시청</h4><p>서울의 중심지입니다.</p></div>'
    });

    // 마커 클릭 시 정보 창 표시
    window.naver.maps.Event.addListener(marker, 'click', () => {
      if (infoWindow.getMap()) {
        infoWindow.close();
      } else {
        infoWindow.open(mapInstance.current, marker);
      }
    });

  }, []);

  // 지도 크기 조정 함수 (창 크기 변경 시 사용)
  useEffect(() => {
    const handleResize = () => {
      if (mapInstance.current) {
        mapInstance.current.autoResize();
      }
    };

    window.addEventListener('resize', handleResize);
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return (
    <div>
      <h1>네이버 지도 예제</h1>
      <div 
        ref={mapElement} 
        style={{ 
          width: '100%', 
          height: '500px',
          border: '1px solid #ddd',
          borderRadius: '8px'
        }}
      />
    </div>
  );
};

export default Map;