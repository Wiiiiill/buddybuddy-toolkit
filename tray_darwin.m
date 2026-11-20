#import <Cocoa/Cocoa.h>

extern void tray_on_show(void);
extern void tray_on_quit(void);

// 类名刻意避开 AppDelegate，避免与 Wails 的同名 ObjC 类在链接时冲突
@interface BuddyTrayController : NSObject
@property (strong) NSStatusItem *statusItem;
@end

@implementation BuddyTrayController

- (void)showWindow:(id)sender {
    tray_on_show();
}

- (void)quitApp:(id)sender {
    tray_on_quit();
}

@end

static BuddyTrayController *gTrayController = nil;

void tray_setup(const char *showText, const char *quitText, const char *tooltip, const void *iconData, int iconLen) {
    NSString *showStr = [NSString stringWithUTF8String:showText ? showText : ""];
    NSString *quitStr = [NSString stringWithUTF8String:quitText ? quitText : ""];
    NSString *tipStr = [NSString stringWithUTF8String:tooltip ? tooltip : ""];
    NSData *iconCopy = [NSData dataWithBytes:iconData length:(NSUInteger)iconLen];

    // Wails 的主线程已经在跑 NSApplication，这里只把创建动作投递到主队列
    dispatch_async(dispatch_get_main_queue(), ^{
        if (gTrayController != nil) {
            return;
        }
        gTrayController = [[BuddyTrayController alloc] init];

        NSStatusItem *item = [[NSStatusBar systemStatusBar]
            statusItemWithLength:NSVariableStatusItemLength];
        gTrayController.statusItem = item;

        NSImage *image = [[NSImage alloc] initWithData:iconCopy];
        if (image != nil) {
            [image setSize:NSMakeSize(18.0, 18.0)];
            item.button.image = image;
        }
        item.button.toolTip = tipStr;

        // macOS 菜单栏惯例：点击图标显示菜单
        NSMenu *menu = [[NSMenu alloc] init];
        NSMenuItem *showItem = [[NSMenuItem alloc] initWithTitle:showStr
                                                         action:@selector(showWindow:)
                                                  keyEquivalent:@""];
        showItem.target = gTrayController;
        [menu addItem:showItem];
        [menu addItem:[NSMenuItem separatorItem]];

        NSMenuItem *quitItem = [[NSMenuItem alloc] initWithTitle:quitStr
                                                         action:@selector(quitApp:)
                                                  keyEquivalent:@""];
        quitItem.target = gTrayController;
        [menu addItem:quitItem];

        item.menu = menu;
    });
}

void tray_remove(void) {
    dispatch_async(dispatch_get_main_queue(), ^{
        if (gTrayController == nil) {
            return;
        }
        if (gTrayController.statusItem != nil) {
            [[NSStatusBar systemStatusBar] removeStatusItem:gTrayController.statusItem];
            gTrayController.statusItem = nil;
        }
        gTrayController = nil;
    });
}
