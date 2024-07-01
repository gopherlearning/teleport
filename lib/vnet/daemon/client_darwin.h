#ifndef client_darwin_h
#define client_darwin_h

#import <Foundation/Foundation.h>
#import "protocol_darwin.h"

typedef struct BundlePathResult {
    const char * bundlePath;
} BundlePathResult;

// BundlePath updates bundlePath of result with the path
// to the bundle directory that contains the current executable.
// bundlePath is an empty string if the bundle details could not be fetched.
// It might return a path even for executables that are not in a bundle.
// In that case, calling codesign --verify on that path will simply return with 1.
void BundlePath(struct BundlePathResult *result);

// Returns the label for the daemon by getting the identifier of the bundle
// this executable is shipped in and appending ".vnetd" to it.
//
// The returned string might be empty if the executable is not in a bundle.
//
// The filename and the value of the Label key in the plist file and the Mach
// service of of the daemon must match the string returned from this function.
NSString * DaemonLabel(void);

// DaemonPlist takes the result of DaemonLabel and appends ".plist" to it
// if not empty.
NSString * DaemonPlist(void);

typedef struct RegisterDaemonResult {
    bool ok;
    // service_status is fetched even if [service registerAndReturnError] fails,
    // for debugging purposes.
    int service_status;
    const char * error_description;
} RegisterDaemonResult;

// RegisterDaemon attempts to register the daemon. After the registration attempt,
// it fetches the daemon status.
// Pretty much a noop if the daemon is already registered and enabled.
void RegisterDaemon(struct RegisterDaemonResult *result);

// DaemonStatus returns the current status of the daemon's service in SMAppService.
// Returns -1 if the given macOS version doesn't support SMAppService.
int DaemonStatus(void);

void OpenSystemSettingsLoginItems(void);

typedef struct StartVnetRequest {
    struct VnetParams * vnet_params;
} StartVnetRequest;

typedef struct StartVnetResponse {
    bool ok;
    const char * error_domain;
    const char * error_description;
} StartVnetResponse;

// StartVnet spawns the daemon process. Only the first call does that,
// subsequent calls are noops. The daemon process exits after the socket file
// in request.vnet_params.socket_path is removed. After that it can be spawned
// again by calling StartVnet.
//
// Blocks until the daemon receives the message or until the client gets
// invalidated.
//
// After calling StartVnet, the caller is expected to call InvalidateDaemonClient
// when a surrounding context in Go gets canceled.
void StartVnet(struct StartVnetRequest *request, struct StartVnetResponse *response);

// InvalidateDaemonClient closes the connection to the daemon and unblocks
// any calls awaiting a reply from the daemon.
void InvalidateDaemonClient(void);

@interface VNEDaemonClient : NSObject
-(void)startVnet:(VnetParams *)vnetParams completion:(void (^)(NSError * error))completion;
// invalidate executes all outstanding reply blocks, error handling blocks,
// and invalidation blocks and forbids from sending or receiving new messages.
-(void)invalidate;
@end

// VNECopyNSString duplicates an NSString into an UTF-8 encoded C string.
// The caller is expected to free the returned pointer.
char *VNECopyNSString(NSString *val);

#endif /* client_darwin_h */
