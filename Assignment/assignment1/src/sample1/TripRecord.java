public class TripRecord {

    private String pickup_DateTime;
    private GPScoord pickup_Location;
    private GPScoord dropoff_Location;
    private double trip_Distance;
    private boolean visited;
    private int cluster_id;

    public TripRecord(String pickup_DateTime, GPScoord pickup_Location, GPScoord dropoff_Location, double trip_Distance){
        this.pickup_DateTime = pickup_DateTime;
        this.pickup_Location = pickup_Location;
        this.dropoff_Location = dropoff_Location;
        this.trip_Distance = trip_Distance;
        this.visited = false;
        this.cluster_id = -1;
    }

    public boolean is_visited(){
        return this.visited;
    }

    public void set_visited(boolean visitNode){
        this.visited = visitNode;
    }

    public int get_cluster_id(){
        return this.cluster_id;
    }

    public void set_cluster_id(int cluster_id){
        this.cluster_id = cluster_id;
    }

    public String pickup_date_time(){
        return this.pickup_DateTime;
    }

    public GPScoord pickup_location(){
        return this.pickup_Location;
    }

    public GPScoord dropoff_location(){
        return this.dropoff_Location;
    }

    public double trip_distance(){
        return this.trip_Distance;
    }

    @Override
    public boolean equals(Object o){
        //Check if Object o is null or not an instance of TripRecord
        if(o == null || !(o instanceof TripRecord)){
            return false;
        }
        //If object being compared to itself, then true
        if(o == this){
            return true;
        }
        
        TripRecord tr = (TripRecord) o;
        
        return 
        this.pickup_date_time() == tr.pickup_date_time() &&
        this.pickup_location().get_latitude() == tr.pickup_location().get_latitude() &&
        this.pickup_location().get_longitude() == tr.pickup_location().get_longitude() &&
        this.dropoff_location().get_latitude() == tr.dropoff_location().get_latitude() &&
        this.dropoff_location().get_longitude() == tr.dropoff_location().get_longitude() &&
        this.trip_distance() == tr.trip_distance();
    }
}
