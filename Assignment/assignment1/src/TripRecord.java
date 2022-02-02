public class TripRecord {

    //attributes
    private String pickup_DateTime;
    private GPScoord pickup_Location;
    private GPScoord dropoff_Location;
    private double trip_distance;


    private Boolean visited; 
    private int cluster_id; 
    //constructor 
    public TripRecord(String pickupDateTime, GPScoord pickupLocation, GPScoord dropoffLocation, double tripDistance){
        this.pickup_DateTime = pickupDateTime; 
        this.pickup_Location = pickupLocation; 
        this.dropoff_Location = dropoffLocation; 
        this.trip_distance = tripDistance; 
        this.visited = false; 
        this.cluster_id = -1; 
    }   


    public void visit(){
       visited = true;  
    }

    public Boolean is_visited(){
        return visited; 
    }

    public int getClusterId(){
        return cluster_id; 
    }

    public String getPickupDateTime(){
        return pickup_DateTime; 
    }

    public GPScoord getPickupLocation(){
        return pickup_Location; 
    }

    public GPScoord getDropoffLocation(){
        return dropoff_Location; 
    }

    public double getTripDistance(){
        return trip_distance; 
    }

}
