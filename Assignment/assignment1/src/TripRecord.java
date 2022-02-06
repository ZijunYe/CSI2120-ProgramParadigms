public class TripRecord {

    //attributes
    private String pickup_DateTime;
    private GPScoord pickup_Location;
    private GPScoord dropoff_Location;
    private double trip_distance;


    private Boolean visited; 
    private int clusterid; 
    private boolean isNoise;

    
    //constructor 
    public TripRecord(String pickupDateTime, GPScoord pickupLocation, GPScoord dropoffLocation, double tripDistance){
        this.pickup_DateTime = pickupDateTime; 
        this.pickup_Location = pickupLocation; 
        this.dropoff_Location = dropoffLocation; 
        this.trip_distance = tripDistance; 
        this.visited = false; 
        this.clusterid = -1;
        this.isNoise = false; 
       
    }   

    public void setNoise(){
        isNoise = true; 

    } 

    public Boolean isNoise(){
        return isNoise; 
    }
    public void visit(){
       visited = true;  
    }

    public Boolean is_visited(){
        return visited; 
    }


    public int getClusterId(){
        return clusterid; 
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


    public void getInCluster(int id) {
        this.clusterid = id; 
    }

    public Boolean equals(TripRecord other){
        Boolean result = (
            this.getDropoffLocation().equals(other.getDropoffLocation()) &&
            this.getPickupDateTime().equals(other.getPickupDateTime())&&
            this.getPickupLocation().equals(other.getPickupLocation()) &&
            this.getTripDistance() == other.getTripDistance()
        );
        return result; 
    }

}
