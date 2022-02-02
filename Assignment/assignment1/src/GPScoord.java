public class GPScoord {
    //two attributes for GPS coordinates 
    private  double longitude; 
    private double latitude; 

    //constructor 
    public GPScoord(double longitude, double latitude){
        this.longitude = longitude; 
        this.latitude = latitude; 
    }

    //getter method 
    public double getLongitude(){
        return longitude; 
    }

    public double getLatitude(){
        return latitude; 
    }

}




