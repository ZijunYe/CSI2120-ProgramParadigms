public class GPScoord{

    private double latitude;
    private double longitude;

    public GPScoord(double latitude, double longitude){
        this.latitude = latitude;
        this.longitude = longitude;
    }

    public double get_latitude(){
        return this.latitude;
    }

    public double get_longitude(){
        return this.longitude;
    }
}